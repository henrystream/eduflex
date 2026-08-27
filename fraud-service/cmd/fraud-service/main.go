package main

import (
	"log"
	"net/http"

	"github.com/henrystream/eduflex/fraud-service/internal/client"
	"github.com/henrystream/eduflex/fraud-service/internal/config"
	apphttp "github.com/henrystream/eduflex/fraud-service/internal/http"
	"github.com/henrystream/eduflex/fraud-service/internal/service"
)

func main() {
	cfg := config.Load()

	studentClient := client.NewStudentClient(cfg.StudentURL)
	schoolClient := client.NewSchoolClient(cfg.SchoolURL)
	financingClient := client.NewFinancingClient(cfg.FinancingURL)
	loanClient := client.NewLoanClient(cfg.LoanURL)
	ledgerClient := client.NewLedgerClient(cfg.LedgerURL)
	eventsClient := client.NewEventsClient(cfg.EventsURL, "fraud-service")
	scoringEngine := service.NewScoringEngine()

	fraudSvc := service.NewFraudService(
		studentClient,
		schoolClient,
		financingClient,
		loanClient,
		ledgerClient,
		eventsClient,
		scoringEngine,
	)

	router := apphttp.NewRouter(fraudSvc)

	log.Printf("fraud-service running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
