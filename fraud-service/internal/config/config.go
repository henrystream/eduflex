package config

import "os"

type Config struct {
	StudentURL   string
	SchoolURL    string
	FinancingURL string
	LoanURL      string
	LedgerURL    string
	EventsURL    string
	FraudURL     string
	Port         string
}

func Load() Config {
	return Config{
		StudentURL:   getEnv("STUDENT_SERVICE_URL", "http://student-service:8080"),
		SchoolURL:    getEnv("SCHOOL_SERVICE_URL", "http://school-service:8080"),
		FinancingURL: getEnv("FINANCING_SERVICE_URL", "http://financing-service:8080"),
		LoanURL:      getEnv("LOAN_SERVICE_URL", "http://loan-service:8080"),
		LedgerURL:    getEnv("LEDGER_SERVICE_URL", "http://ledger-service:8080"),
		EventsURL:    getEnv("EVENTS_SERVICE_URL", "http://events-service:8080"),
		FraudURL:     getEnv("FRAUD_SERVICE_URL", "http://fraud-service:8080"),
		Port:         getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
