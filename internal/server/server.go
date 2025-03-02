package server

import "github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func (s *Server) Run() error {
	return nil
}
