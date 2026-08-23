package main

import (
	"context"
	"log"
	"net/http"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/config"
	apphttp "github.com/henrystream/eduflex/financing-service/internal/http"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/henrystream/eduflex/financing-service/internal/service"

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
	agreementRepo := repository.NewAgreementRepository(queries)
	installmentRepo := repository.NewInstallmentRepository(queries)
	installmentSvc := service.NewInstallmentService(installmentRepo)
	agreementSvc := service.NewAgreementService(agreementRepo, installmentSvc)

	router := apphttp.NewRouter(agreementSvc, installmentSvc)

	log.Printf("student-service listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
