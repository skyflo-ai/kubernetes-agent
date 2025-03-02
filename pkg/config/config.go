package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port    int           `mapstructure:"port"`
		Host    string        `mapstructure:"host"`
		Timeout time.Duration `mapstructure:"timeout"`
	}

	Kubernetes struct {
		PollInterval time.Duration `mapstructure:"poll_interval"`
		ClusterName  string        `mapstructure:"cluster_name"`
	}

	API struct {
		Key    string `mapstructure:"key"`
		Server string `mapstructure:"server"`
	}
}

func Load() (*Config, error) {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.timeout", time.Second*30)
	viper.SetDefault("kubernetes.poll_interval", time.Second*30)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SKYFLO")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
