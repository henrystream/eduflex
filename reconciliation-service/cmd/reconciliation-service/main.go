package main

import (
	"net/http"

	"github.com/henrystream/eduflex/reconciliation-service/internal/client"
	"github.com/henrystream/eduflex/reconciliation-service/internal/config"
	apphttp "github.com/henrystream/eduflex/reconciliation-service/internal/http"
	"github.com/henrystream/eduflex/reconciliation-service/internal/service"
)

func main() {
	cfg := config.Load()

	ledgerClient := client.NewLedgerClient(cfg.LedgerURL)
	eventsClient := client.NewEventsClient(cfg.EventsURL, "reconciliation-service")
	matcher := service.NewMatchingEngine()

	svc := service.NewReconciliationService(ledgerClient, eventsClient, matcher)
	router := apphttp.NewRouter(svc)

	http.ListenAndServe(":"+cfg.Port, router)
}
