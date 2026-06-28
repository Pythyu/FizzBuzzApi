package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

const defaultServerPort = 8080
const defaultServerWriteTimeout = 5
const defaultServerReadTimeout = 3
const defaultServerIdleTimeout = 5

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE"`
	Debug        bool          `env:"SERVER_DEBUG"`
}

func NewServerConfig() *ServerConfig {
	config := &ServerConfig{
		Port:         defaultServerPort,
		TimeoutRead:  defaultServerReadTimeout * time.Second,
		TimeoutWrite: defaultServerWriteTimeout * time.Second,
		TimeoutIdle:  defaultServerIdleTimeout * time.Second,
		Debug:        false,
	}
	if err := envdecode.Decode(config); err != nil {
		log.Fatalf("Failed to decode the server config: %s", err)
	}

	return config
}
