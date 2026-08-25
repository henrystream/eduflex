package config

import "os"

type Config struct {
	EventsURL string
	EmailFrom string
	SMSFrom   string
	Interval  int
	Port      string
}

func Load() Config {
	return Config{
		EventsURL: getEnv("EVENTS_SERVICE_URL", "http://events-service:8080"),
		EmailFrom: getEnv("EMAIL_FROM", "noreply@eduflex.com"),
		SMSFrom:   getEnv("SMS_FROM", "Eduflex"),
		Interval:  5,
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
