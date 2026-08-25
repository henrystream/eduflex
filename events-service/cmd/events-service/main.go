package main

import (
	"context"
	"log"
	"net/http"

	db "github.com/henrystream/eduflex/events-service/db/sqlc"
	"github.com/henrystream/eduflex/events-service/internal/config"
	apphttp "github.com/henrystream/eduflex/events-service/internal/http"
	"github.com/henrystream/eduflex/events-service/internal/repository"
	"github.com/henrystream/eduflex/events-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db open error: %v", err)
	}
	defer conn.Close()

	queries := db.New(conn)
	repo := repository.NewEventRepository(queries)
	svc := service.NewEventService(repo)
	router := apphttp.NewRouter(svc)

	log.Printf("events-service running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
