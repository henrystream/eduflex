package main

import (
	"context"
	"log"
	"net/http"

	db "github.com/henrystream/eduflex/auth-service/db/sqlc"
	"github.com/henrystream/eduflex/auth-service/internal/config"
	apphttp "github.com/henrystream/eduflex/auth-service/internal/http"
	"github.com/henrystream/eduflex/auth-service/internal/repository"
	"github.com/henrystream/eduflex/auth-service/internal/service"
	"github.com/henrystream/eduflex/auth-service/internal/token"
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
	repo := repository.NewUserRepository(queries)
	jwtMaker := token.NewJWTMaker(cfg.JWTSecret)
	svc := service.NewAuthService(repo, jwtMaker)
	router := apphttp.NewRouter(svc)

	log.Printf("auth-service running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
