package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	apphttp "github.com/henrystream/eduflex/student-service/internal/http"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/config"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/henrystream/eduflex/student-service/internal/service"
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
	studentRepo := repository.NewStudentRepository(queries)
	studentSvc := service.NewStudentService(studentRepo)

	enrollRepo := repository.NewEnrollmentRepository(queries)
	enrollSvc := service.NewEnrollmentService(enrollRepo)

	payRepo := repository.NewPaymentRepository(queries)
	paySvc := service.NewPaymentService(payRepo, ledger)

	router := apphttp.NewRouter(studentSvc, enrollSvc, paySvc)

	log.Printf("student-service listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
