package main

import (
	"log"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/internal/metrics"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	m, err := metrics.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create metrics collector: %v", err)
	}

	if err := m.Run(); err != nil {
		log.Fatalf("Metrics collector failed: %v", err)
	}
}
