package main

import (
	"context"
	"log"
	"net/http"
	"os"

	db "github.com/henrystream/eduflex/disbursement-service/db/sqlc"
	"github.com/henrystream/eduflex/disbursement-service/internal/config"
	apphttp "github.com/henrystream/eduflex/disbursement-service/internal/http"
	"github.com/henrystream/eduflex/disbursement-service/internal/ledger"
	"github.com/henrystream/eduflex/disbursement-service/internal/repository"
	"github.com/henrystream/eduflex/disbursement-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer conn.Close()

	queries := db.New(conn)
	repo := repository.NewDisbursementRepository(queries)

	ledgerURL := os.Getenv("LEDGER_URL")
	if ledgerURL == "" {
		ledgerURL = "http://ledger-service:8080"
	}
	ledgerClient := ledger.NewLedgerClient(ledgerURL)

	svc := service.NewDisbursementService(repo, ledgerClient)
	router := apphttp.NewRouter(svc)

	log.Printf("disbursement-service running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
