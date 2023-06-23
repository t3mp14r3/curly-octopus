package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Config struct {
    ServerConfig    ServerConfig
}

type ServerConfig struct {
    Port            string `env:"SERVER_PORT"`
}

func New() *Config {
    var config Config

    if err := env.Parse(&config.ServerConfig); err != nil {
        log.Fatalf("failed to parse server config: %v", err)
	}

	return &config
}

