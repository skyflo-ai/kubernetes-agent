package types

import (
	"time"
)

// ResourceType represents the type of Kubernetes resource
type ResourceType string

const (
	TypeNode        ResourceType = "node"
	TypeNamespace   ResourceType = "namespace"
	TypeIngress     ResourceType = "ingress"
	TypeService     ResourceType = "service"
	TypeDeployment  ResourceType = "deployment"
	TypeStatefulSet ResourceType = "statefulset"
	TypePod         ResourceType = "pod"
	TypeConfigMap   ResourceType = "configmap"
	TypeSecret      ResourceType = "secret"
)

// EventType represents the type of event
type EventType string

const (
	EventTypeInitial EventType = "INITIAL"
	EventTypeAdd     EventType = "ADD"
	EventTypeUpdate  EventType = "UPDATE"
	EventTypeDelete  EventType = "DELETE"
)

// ResourceEvent represents an event for a Kubernetes resource
type ResourceEvent struct {
	ClusterName  string            `json:"cluster_name"`
	ResourceType ResourceType      `json:"resource_type"`
	EventType    EventType         `json:"event_type"`
	Timestamp    time.Time         `json:"timestamp"`
	Payload      interface{}       `json:"payload"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// ResourceMetadata contains common metadata for resources
type ResourceMetadata struct {
	Name            string
	Namespace       string
	ResourceVersion string
	Labels          map[string]string
	Annotations     map[string]string
}
