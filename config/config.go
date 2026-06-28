package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT" envDefault:"8080"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ" envDefault:"3s"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE" envDefault:"5s"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE" envDefault:"5s"`
	Debug        bool          `env:"SERVER_DEBUG" envDefault:"false"`
}

func NewServerConfig() *ServerConfig {
	config := &ServerConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Failed to decode the server config: %v", err)
	}

	return config
}
