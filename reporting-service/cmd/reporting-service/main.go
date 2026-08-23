package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henrystream/eduflex/reporting-service/internal/client"
	"github.com/henrystream/eduflex/reporting-service/internal/config"
	httphandlers "github.com/henrystream/eduflex/reporting-service/internal/http"
	"github.com/henrystream/eduflex/reporting-service/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize HTTP clients
	ledgerClient := client.NewLedgerClient(cfg.LedgerURL)
	schoolClient := client.NewSchoolClient(cfg.SchoolURL)
	studentClient := client.NewStudentClient(cfg.StudentURL)
	financingClient := client.NewFinancingClient(cfg.FinancingURL)

	// Initialize services
	schoolReportService := service.NewSchoolReportService(ledgerClient, studentClient, schoolClient)
	studentReportService := service.NewStudentReportService(ledgerClient, financingClient, studentClient)
	financialReportService := service.NewFinancialReportService(ledgerClient)

	// Create router
	router := httphandlers.NewRouter(schoolReportService, studentReportService, financialReportService)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting reporting service on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
