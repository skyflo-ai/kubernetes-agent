package main

import (
	"context"
	"log"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/internal/watcher"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	w, err := watcher.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	if err := w.Run(context.Background()); err != nil {
		log.Fatalf("Watcher failed: %v", err)
	}
}
