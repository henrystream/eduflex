package main

import (
	"log"
	"net/http"

	"github.com/henrystream/eduflex/api-gateway/internal/config"
	apphttp "github.com/henrystream/eduflex/api-gateway/internal/http"
)

func main() {
	cfg := config.Load()
	router := apphttp.NewRouter(cfg)

	log.Printf("API Gateway running on :%s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, router)
}
