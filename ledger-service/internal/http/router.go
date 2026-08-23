package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/ledger-service/internal/service"
)

func NewRouter(svc *service.LedgerService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewLedgerHandler(svc)

	r.Route("/ledger", func(r chi.Router) {
		r.Post("/", h.CreateEntry)
		r.Get("/", h.ListAll)
		r.Get("/by-event", h.ListByEvent)
		r.Get("/by-service", h.ListByService)
	})

	return r
}
