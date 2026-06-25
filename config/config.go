package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

func NewServerConfig() *ServerConfig {
	var config ServerConfig
	if err := envdecode.StrictDecode(&config); err != nil {
		log.Fatalf("Failed to decode the server config: %s", err)
	}

	return &config
}
