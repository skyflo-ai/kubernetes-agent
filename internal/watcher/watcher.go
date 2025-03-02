package watcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/sender"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/types"
)

type Watcher struct {
	cfg             *config.Config
	client          kubernetes.Interface
	sender          *sender.Sender
	informerFactory informers.SharedInformerFactory
	factory         *resourceWatcherFactory
	healthy         bool
	mu              sync.RWMutex
}

func New(cfg *config.Config) (*Watcher, error) {
	k8sConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create in-cluster config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Hour*24)
	sender := sender.New(cfg)

	return &Watcher{
		cfg:             cfg,
		client:          clientset,
		sender:          sender,
		informerFactory: informerFactory,
		factory:         newResourceWatcherFactory(informerFactory, sender, cfg.Kubernetes.ClusterName),
	}, nil
}

func (w *Watcher) Run(ctx context.Context) error {
	defer func() {
		w.mu.Lock()
		w.healthy = false
		w.mu.Unlock()
	}()

	w.mu.Lock()
	w.healthy = true
	w.mu.Unlock()

	// Start informer factory
	w.informerFactory.Start(ctx.Done())

	// Wait for initial sync
	cachesSynced := w.informerFactory.WaitForCacheSync(ctx.Done())
	for _, synced := range cachesSynced {
		if !synced {
			return fmt.Errorf("failed to sync caches")
		}
	}

	// Initial resource crawl
	if err := w.initialCrawl(ctx); err != nil {
		return fmt.Errorf("initial crawl failed: %w", err)
	}

	// Setup watchers
	w.setupWatchers()

	<-ctx.Done()
	return ctx.Err()
}

func (w *Watcher) setupWatchers() {
	// Setup informers with the factory's event handlers
	w.informerFactory.Core().V1().Nodes().Informer().AddEventHandler(
		w.factory.createEventHandlers(types.TypeNode))
	// Add other resources similarly...
}

func (w *Watcher) IsHealthy() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.healthy
}

func (w *Watcher) initialCrawl(ctx context.Context) error {
	// Get initial state of all resources
	if err := w.crawlNodes(ctx); err != nil {
		return err
	}
	if err := w.crawlNamespaces(ctx); err != nil {
		return err
	}
	if err := w.crawlIngresses(ctx); err != nil {
		return err
	}
	if err := w.crawlServices(ctx); err != nil {
		return err
	}
	if err := w.crawlDeployments(ctx); err != nil {
		return err
	}
	if err := w.crawlStatefulSets(ctx); err != nil {
		return err
	}
	if err := w.crawlPods(ctx); err != nil {
		return err
	}
	if err := w.crawlConfigMaps(ctx); err != nil {
		return err
	}
	if err := w.crawlSecrets(ctx); err != nil {
		return err
	}
	return nil
}

func (w *Watcher) watchResources(ctx context.Context) error {
	errCh := make(chan error, 9) // One for each resource type

	// Start watchers for each resource type
	go w.watchNodes(ctx, errCh)
	go w.watchNamespaces(ctx, errCh)
	go w.watchIngresses(ctx, errCh)
	go w.watchServices(ctx, errCh)
	go w.watchDeployments(ctx, errCh)
	go w.watchStatefulSets(ctx, errCh)
	go w.watchPods(ctx, errCh)
	go w.watchConfigMaps(ctx, errCh)
	go w.watchSecrets(ctx, errCh)

	// Wait for any error or context cancellation
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
