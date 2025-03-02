package main

import (
	"log"

	"github.com/KaranJagtiani/skyflo-kubernetes-agent/internal/server"
	"github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	s, err := server.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
