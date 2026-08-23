package main

import (
	"context"
	"log"
	"net/http"

	db "github.com/henrystream/eduflex/ledger-service/db/sqlc"
	"github.com/henrystream/eduflex/ledger-service/internal/config"
	apphttp "github.com/henrystream/eduflex/ledger-service/internal/http"
	"github.com/henrystream/eduflex/ledger-service/internal/repository"
	"github.com/henrystream/eduflex/ledger-service/internal/service"
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

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	queries := db.New(conn)
	repo := repository.NewLedgerRepository(queries)
	svc := service.NewLedgerService(repo)
	router := apphttp.NewRouter(svc)

	log.Printf("ledger-service listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
