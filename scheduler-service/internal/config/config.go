package config

import "os"

type Config struct {
	FinancingURL string
	StudentURL   string
	LoanURL      string
	EventsURL    string
	Interval     int
}

func Load() Config {
	return Config{
		FinancingURL: getEnv("FINANCING_SERVICE_URL", "http://financing-service:8080"),
		StudentURL:   getEnv("STUDENT_SERVICE_URL", "http://student-service:8080"),
		LoanURL:      getEnv("LOAN_SERVICE_URL", "http://loan-service:8080"),
		EventsURL:    getEnv("EVENTS_SERVICE_URL", "http://events-service:8080"),
		Interval:     10, // run every 10 seconds
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
