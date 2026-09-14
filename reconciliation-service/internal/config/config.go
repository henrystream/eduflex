package config

import "os"

type Config struct {
	LedgerURL       string
	LoanURL         string
	DisbursementURL string
	FinancingURL    string
	EventsURL       string
	Port            string
}

func Load() Config {
	return Config{
		LedgerURL:       getEnv("LEDGER_SERVICE_URL", "http://ledger-service:8080"),
		LoanURL:         getEnv("LOAN_SERVICE_URL", "http://loan-service:8080"),
		DisbursementURL: getEnv("DISBURSEMENT_SERVICE_URL", "http://disbursement-service:8080"),
		FinancingURL:    getEnv("FINANCING_SERVICE_URL", "http://financing-service:8080"),
		EventsURL:       getEnv("EVENTS_SERVICE_URL", "http://events-service:8080"),
		Port:            getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
