package events

import "sync"

type EventType string

const (
	PodCreated    EventType = "POD_CREATED"
	PodDeleted    EventType = "POD_DELETED"
	MetricsUpdate EventType = "METRICS_UPDATE"
)

type Event struct {
	Type    EventType
	Payload interface{}
}

type EventBus struct {
	subscribers map[EventType][]chan Event
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]chan Event),
	}
}
