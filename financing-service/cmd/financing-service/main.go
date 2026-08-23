package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/config"
	apphttp "github.com/henrystream/eduflex/financing-service/internal/http"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/henrystream/eduflex/financing-service/internal/service"

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
	agreementRepo := repository.NewAgreementRepository(queries)
	installmentRepo := repository.NewInstallmentRepository(queries)
	installmentSvc := service.NewInstallmentService(installmentRepo, ledger)
	agreementSvc := service.NewAgreementService(agreementRepo, installmentSvc, ledger)

	router := apphttp.NewRouter(agreementSvc, installmentSvc)

	log.Printf("financing-service listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
