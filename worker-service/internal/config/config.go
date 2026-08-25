package config

import "os"

type Config struct {
	EventsURL string
	Interval  int // seconds
}

func Load() Config {
	cfg := Config{
		EventsURL: os.Getenv("EVENTS_SERVICE_URL"),
		Interval:  5, //os.Getenv("EVENT_INTERVAL"),
	}

	if cfg.EventsURL == "" {
		cfg.EventsURL = "http://events-service:8086"
	}

	return cfg

}
