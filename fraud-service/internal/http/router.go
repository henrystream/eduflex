package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/fraud-service/internal/service"
)

func NewRouter(svc *service.FraudService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewFraudHandler(svc)

	r.Get("/fraud/check-agreement", h.CheckAgreement)

	return r
}
