package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	apphttp "github.com/henrystream/eduflex/loan-service/internal/http"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/config"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/henrystream/eduflex/loan-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type ledgerClient struct {
	baseURL string
}

func (c ledgerClient) CreateEntry(req service.LedgerEntryRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.baseURL+"/ledger", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(resp.Status)
	}
	return nil
}

func main() {
	cfg := config.Load()

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer conn.Close()

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	ledgerURL := os.Getenv("LEDGER_URL")
	if ledgerURL == "" {
		ledgerURL = "http://ledger-service:8080"
	}
	ledger := ledgerClient{baseURL: ledgerURL}

	queries := db.New(conn)

	bankRepo := repository.NewBankRepository(queries)
	facilityRepo := repository.NewFacilityRepository(queries)
	drawdownRepo := repository.NewDrawdownRepository(queries)
	repaymentRepo := repository.NewRepaymentRepository(queries)

	bankSvc := service.NewBankService(bankRepo)
	facilitySvc := service.NewFacilityService(facilityRepo)
	drawdownSvc := service.NewDrawdownService(drawdownRepo, ledger)
	repaymentSvc := service.NewRepaymentService(repaymentRepo, ledger)

	router := apphttp.NewRouter(bankSvc, facilitySvc, drawdownSvc, repaymentSvc)

	log.Printf("loan-service listening on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
