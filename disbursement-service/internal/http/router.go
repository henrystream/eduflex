package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/disbursement-service/internal/service"
)

func NewRouter(svc *service.DisbursementService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewDisbursementHandler(svc)

	r.Route("/disbursements", func(r chi.Router) {
		r.Post("/", h.CreateDisbursement)
		r.Get("/", h.ListDisbursementsBySchool)
	})

	return r
}
