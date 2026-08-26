package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/reconciliation-service/internal/service"
)

func NewRouter(svc *service.ReconciliationService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewReconciliationHandler(svc)

	r.Get("/reconcile", h.Reconcile)

	return r
}
