package main

import (
	"context"
	"log"
	"net/http"

	apphttp "github.com/henrystream/eduflex/student-service/internal/http"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/config"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/henrystream/eduflex/student-service/internal/service"
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
	studentRepo := repository.NewStudentRepository(queries)
	studentSvc := service.NewStudentService(studentRepo)

	enrollRepo := repository.NewEnrollmentRepository(queries)
	enrollSvc := service.NewEnrollmentService(enrollRepo)

	payRepo := repository.NewPaymentRepository(queries)
	paySvc := service.NewPaymentService(payRepo)

	router := apphttp.NewRouter(studentSvc, enrollSvc, paySvc)

	log.Printf("student-service listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
