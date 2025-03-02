package watcher

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/types"
)

// crawlNodes gets the initial state of nodes
func (w *Watcher) crawlNodes(ctx context.Context) error {
	nodes, err := w.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	// Send initial state to server
	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeNode,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      nodes.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send nodes data: %w", err)
	}

	return nil
}

// crawlNamespaces gets the initial state of namespaces
func (w *Watcher) crawlNamespaces(ctx context.Context) error {
	namespaces, err := w.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeNamespace,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      namespaces.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send namespaces data: %w", err)
	}

	return nil
}

// crawlIngresses gets the initial state of ingresses
func (w *Watcher) crawlIngresses(ctx context.Context) error {
	ingresses, err := w.client.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeIngress,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      ingresses.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send ingresses data: %w", err)
	}

	return nil
}

// crawlServices gets the initial state of services
func (w *Watcher) crawlServices(ctx context.Context) error {
	services, err := w.client.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeService,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      services.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send services data: %w", err)
	}

	return nil
}

// crawlDeployments gets the initial state of deployments
func (w *Watcher) crawlDeployments(ctx context.Context) error {
	deployments, err := w.client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeDeployment,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      deployments.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send deployments data: %w", err)
	}

	return nil
}

// crawlStatefulSets gets the initial state of statefulsets
func (w *Watcher) crawlStatefulSets(ctx context.Context) error {
	statefulsets, err := w.client.AppsV1().StatefulSets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeStatefulSet,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      statefulsets.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send statefulsets data: %w", err)
	}

	return nil
}

// crawlPods gets the initial state of pods
func (w *Watcher) crawlPods(ctx context.Context) error {
	pods, err := w.client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypePod,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      pods.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send pods data: %w", err)
	}

	return nil
}

// crawlConfigMaps gets the initial state of configmaps
func (w *Watcher) crawlConfigMaps(ctx context.Context) error {
	configmaps, err := w.client.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeConfigMap,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      configmaps.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send configmaps data: %w", err)
	}

	return nil
}

// crawlSecrets gets the initial state of secrets
func (w *Watcher) crawlSecrets(ctx context.Context) error {
	secrets, err := w.client.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	err = w.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  w.cfg.Kubernetes.ClusterName,
		ResourceType: types.TypeSecret,
		EventType:    types.EventTypeInitial,
		Timestamp:    time.Now(),
		Payload:      secrets.Items,
	})
	if err != nil {
		return fmt.Errorf("failed to send secrets data: %w", err)
	}

	return nil
}

// watchNodes sets up a watcher for nodes
func (w *Watcher) watchNodes(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().Nodes().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNode,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send node add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNode,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send node update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNode,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send node delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchNamespaces sets up a watcher for namespaces
func (w *Watcher) watchNamespaces(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().Namespaces().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNamespace,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send namespace add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNamespace,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send namespace update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeNamespace,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send namespace delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchIngresses sets up a watcher for ingresses
func (w *Watcher) watchIngresses(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Networking().V1().Ingresses().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeIngress,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send ingress add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeIngress,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send ingress update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeIngress,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send ingress delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchServices sets up a watcher for services
func (w *Watcher) watchServices(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().Services().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeService,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send service add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeService,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send service update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeService,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send service delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchDeployments sets up a watcher for deployments
func (w *Watcher) watchDeployments(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Apps().V1().Deployments().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeDeployment,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send deployment add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeDeployment,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send deployment update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeDeployment,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send deployment delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchStatefulSets sets up a watcher for statefulsets
func (w *Watcher) watchStatefulSets(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Apps().V1().StatefulSets().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeStatefulSet,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send statefulset add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeStatefulSet,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send statefulset update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeStatefulSet,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send statefulset delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchPods sets up a watcher for pods
func (w *Watcher) watchPods(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypePod,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send pod add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypePod,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send pod update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypePod,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send pod delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchConfigMaps sets up a watcher for configmaps
func (w *Watcher) watchConfigMaps(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().ConfigMaps().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeConfigMap,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send configmap add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeConfigMap,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send configmap update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeConfigMap,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send configmap delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}

// watchSecrets sets up a watcher for secrets
func (w *Watcher) watchSecrets(ctx context.Context, errCh chan<- error) {
	factory := informers.NewSharedInformerFactory(w.client, time.Hour*24)
	informer := factory.Core().V1().Secrets().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeSecret,
				EventType:    types.EventTypeAdd,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send secret add event: %w", err)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeSecret,
				EventType:    types.EventTypeUpdate,
				Timestamp:    time.Now(),
				Payload:      new,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send secret update event: %w", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if err := w.sender.SendResourceEvent(ctx, types.ResourceEvent{
				ClusterName:  w.cfg.Kubernetes.ClusterName,
				ResourceType: types.TypeSecret,
				EventType:    types.EventTypeDelete,
				Timestamp:    time.Now(),
				Payload:      obj,
			}); err != nil {
				errCh <- fmt.Errorf("failed to send secret delete event: %w", err)
			}
		},
	})

	informer.Run(ctx.Done())
}
