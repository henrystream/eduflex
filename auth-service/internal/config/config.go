package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

func Load() Config {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("SECRET"),
	}

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = "postgres://postgres:password@authdb:5435/authdb?sslmode=disable"
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "password"
	}

	return cfg
}
