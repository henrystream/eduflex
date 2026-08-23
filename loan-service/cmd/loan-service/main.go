package main

import (
	"context"
	"log"
	"net/http"

	apphttp "github.com/henrystream/eduflex/loan-service/internal/http"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/config"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/henrystream/eduflex/loan-service/internal/service"
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

	bankRepo := repository.NewBankRepository(queries)
	facilityRepo := repository.NewFacilityRepository(queries)
	drawdownRepo := repository.NewDrawdownRepository(queries)
	repaymentRepo := repository.NewRepaymentRepository(queries)

	bankSvc := service.NewBankService(bankRepo)
	facilitySvc := service.NewFacilityService(facilityRepo)
	drawdownSvc := service.NewDrawdownService(drawdownRepo)
	repaymentSvc := service.NewRepaymentService(repaymentRepo)

	router := apphttp.NewRouter(bankSvc, facilitySvc, drawdownSvc, repaymentSvc)

	log.Printf("loan-service listening on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
