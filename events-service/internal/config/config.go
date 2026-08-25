package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() Config {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = "postgres://postgres:postgres@event-db:5432/eventdb?sslmode=disable"
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
