package config

import "os"

type Config struct {
	LedgerURL       string
	FinancingURL    string
	LoanURL         string
	StudentURL      string
	SchoolURL       string
	DisbursementURL string
	Port            string
}

func Load() Config {
	return Config{
		LedgerURL:       getEnv("LEDGER_SERVICE_URL", "http://ledger-service:8080"),
		FinancingURL:    getEnv("FINANCING_SERVICE_URL", "http://financing-service:8080"),
		LoanURL:         getEnv("LOAN_SERVICE_URL", "http://loan-service:8080"),
		StudentURL:      getEnv("STUDENT_SERVICE_URL", "http://student-service:8080"),
		SchoolURL:       getEnv("SCHOOL_SERVICE_URL", "http://school-service:8080"),
		DisbursementURL: getEnv("DISBURSEMENT_SERVICE_URL", "http://disbursement-service:8080"),
		Port:            getEnv("PORT", "8085"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
