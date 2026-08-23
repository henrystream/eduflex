package http

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/henrystream/eduflex/financing-service/internal/service"
)

func NewRouter(agreementsvc *service.AgreementService, installmentsvc *service.InstallmentService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ah := NewAgreementHandler(agreementsvc)
	ih := NewInstallmentHandler(installmentsvc)

	r.Route("/agreements", func(r chi.Router) {
		r.Post("/", ah.CreateAgreement)
	})

	r.Route("/installments", func(r chi.Router) {
		r.Get("/", ih.ListInstallments)
	})

	return r
}
