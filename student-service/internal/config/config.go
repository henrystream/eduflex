package config

import "os"

type Config struct {
	DatabaseURL string
	EventsURL   string
	Port        string
}

func Load() Config {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		EventsURL:   os.Getenv("EVENTS_URL"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = "postgres://postgres:postgres@student-db:5432/studentdb?sslmode=disable"
	}
	if cfg.EventsURL == "" {
		cfg.EventsURL = "http://events-service:8080"
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg

}
