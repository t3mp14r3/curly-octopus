package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Config struct {
    PostgresConfig      PostgresConfig
    ServerConfig        ServerConfig
}

type PostgresConfig struct {
    Host        string `env:"POSTGRES_HOST"`
    Port        int    `env:"POSTGRES_PORT"`
    Name        string `env:"POSTGRES_NAME"`
    User        string `env:"POSTGRES_USER"`
    Password    string `env:"POSTGRES_PASS"`
}

type ServerConfig struct {
    Addr        string `env:"SERVER_ADDR"`
}

func New() *Config {
    var config Config

	if err := env.Parse(&config.PostgresConfig); err != nil {
        log.Fatalf("failed to parse postgres config: %v", err)
	}
	
    if err := env.Parse(&config.ServerConfig); err != nil {
        log.Fatalf("failed to parse server config: %v", err)
	}

	return &config
}

