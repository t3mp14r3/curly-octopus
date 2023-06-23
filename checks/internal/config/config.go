package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Config struct {
    ServerConfig    ServerConfig
}

type ServerConfig struct {
    Addr            string `env:"SERVER_ADDR"`
    FontPath        string `env:"SERVER_FONT_PATH"`
    TemplatePath    string `env:"SERVER_TEMPLATE_PATH"`
    StoragePath     string `env:"SERVER_STORAGE_PATH"`
}

func New() *Config {
    var config Config

    if err := env.Parse(&config.ServerConfig); err != nil {
        log.Fatalf("failed to parse server config: %v", err)
	}

	return &config
}

