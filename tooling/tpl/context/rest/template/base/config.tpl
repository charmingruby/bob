package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	ServerPort       string `env:"SERVER_PORT,required"`
}

func New() (Config, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return Config{}, err
	}

	cfg := Config{
		ServerConfig: serverConfig{
			Port: environment.ServerPort,
		},
	}

	return cfg, nil
}

type Config struct {
	ServerConfig   serverConfig
}

type serverConfig struct {
	Host string
	Port string
}
