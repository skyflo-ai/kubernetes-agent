package sender

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/types"
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

// Sender handles communication with the parent server
type Sender struct {
	cfg        *config.Config
	httpClient *http.Client
}

// New creates a new Sender instance
func New(cfg *config.Config) *Sender {
	return &Sender{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.Server.Timeout,
		},
	}
}

// SendResourceEvent sends a resource event to the parent server
func (s *Sender) SendResourceEvent(ctx context.Context, event types.ResourceEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/v1/resources", s.cfg.API.Server),
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", s.cfg.API.Key)
	req.Header.Set("User-Agent", "skyflo-kubernetes-agent")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("server returned error status: %d", resp.StatusCode)
	}

	return nil
}
