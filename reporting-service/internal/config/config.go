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
	cfg := Config{
		LedgerURL:       os.Getenv("LEDGER_SERVICE_URL"),
		FinancingURL:    os.Getenv("FINANCING_SERVICE_URL"),
		LoanURL:         os.Getenv("LOAN_SERVICE_URL"),
		StudentURL:      os.Getenv("STUDENT_SERVICE_URL"),
		SchoolURL:       os.Getenv("SCHOOL_SERVICE_URL"),
		DisbursementURL: os.Getenv("DISBURSEMENT_SERVICE_URL"),
		Port:            os.Getenv("PORT"),
	}
	if cfg.LedgerURL == "" {
		cfg.LedgerURL = "http://ledger-service:8080"
	}
	if cfg.FinancingURL == "" {
		cfg.FinancingURL = "http://financing-service:8080"
	}
	if cfg.LoanURL == "" {
		cfg.LoanURL = "http://loan-service:8080"
	}
	if cfg.StudentURL == "" {
		cfg.StudentURL = "http://student-service:8080"
	}
	if cfg.SchoolURL == "" {
		cfg.SchoolURL = "http://school-service:8080"
	}
	if cfg.DisbursementURL == "" {
		cfg.DisbursementURL = "http://disbursement-service:8080"
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg

}
