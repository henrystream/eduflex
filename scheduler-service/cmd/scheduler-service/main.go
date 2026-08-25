package main

import (
	"github.com/henrystream/eduflex/scheduler-service/internal/client"
	"github.com/henrystream/eduflex/scheduler-service/internal/config"
	"github.com/henrystream/eduflex/scheduler-service/internal/scheduler"
	"github.com/henrystream/eduflex/scheduler-service/internal/service"
)

func main() {
	cfg := config.Load()

	financingClient := client.NewFinancingClient(cfg.FinancingURL)
	eventsClient := client.NewEventsClient(cfg.EventsURL, "scheduler-service")

	overdueSvc := service.NewOverdueService(financingClient, eventsClient)
	generatorSvc := service.NewInstallmentGenerationService(financingClient, eventsClient)

	sched := scheduler.NewScheduler(overdueSvc, generatorSvc, cfg.Interval)

	// In production, fetch financing IDs from Financing Service
	financingIDs := []string{
		"fin-001",
		"fin-002",
		"fin-003",
	}

	sched.Start(financingIDs)
}
