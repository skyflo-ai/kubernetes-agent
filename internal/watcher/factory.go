package watcher

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/sender"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/types"
)

type resourceWatcherFactory struct {
	informerFactory informers.SharedInformerFactory
	sender          *sender.Sender
	clusterName     string
}

func newResourceWatcherFactory(factory informers.SharedInformerFactory, sender *sender.Sender, clusterName string) *resourceWatcherFactory {
	return &resourceWatcherFactory{
		informerFactory: factory,
		sender:          sender,
		clusterName:     clusterName,
	}
}

func (f *resourceWatcherFactory) createEventHandlers(resourceType types.ResourceType) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			f.handleResourceEvent(context.Background(), obj, resourceType, types.EventTypeAdd)
		},
		UpdateFunc: func(old, new interface{}) {
			f.handleResourceEvent(context.Background(), new, resourceType, types.EventTypeUpdate)
		},
		DeleteFunc: func(obj interface{}) {
			f.handleResourceEvent(context.Background(), obj, resourceType, types.EventTypeDelete)
		},
	}
}

func (f *resourceWatcherFactory) handleResourceEvent(ctx context.Context, obj interface{}, resourceType types.ResourceType, eventType types.EventType) {
	if err := f.sender.SendResourceEvent(ctx, types.ResourceEvent{
		ClusterName:  f.clusterName,
		ResourceType: resourceType,
		EventType:    eventType,
		Timestamp:    time.Now(),
		Payload:      obj,
	}); err != nil {
		// Use structured logging here
		fmt.Printf("failed to send %s %s event: %v\n", resourceType, eventType, err)
	}
}
