package metrics

import "github.com/KaranJagtiani/skyflo-kubernetes-agent/pkg/config"

type Metrics struct {
	cfg *config.Config
}

func New(cfg *config.Config) (*Metrics, error) {
	return &Metrics{cfg: cfg}, nil
}

func (m *Metrics) Run() error {
	return nil
}
