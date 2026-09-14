package http

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/henrystream/eduflex/financing-service/internal/service"
)

func NewRouter(agreementsvc *service.AgreementService, installmentsvc *service.InstallmentService, paymentsvc *service.PaymentService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	agreementsHandler := NewAgreementHandler(agreementsvc)
	r.Route("/agreements", func(r chi.Router) {
		r.Post("/", agreementsHandler.CreateAgreement)
		r.Get("/", agreementsHandler.ListAgreementsByStudent)
	})
	installmentHandler := NewInstallmentHandler(installmentsvc)
	r.Route("/installments", func(r chi.Router) {
		r.Get("/", installmentHandler.ListInstallments)
	})

	paymentHandler := NewPaymentHandler(paymentsvc)
	r.Route("/payments", func(r chi.Router) {
		r.Post("/", paymentHandler.CreatePayment)
		r.Get("/", paymentHandler.ListPaymentsByStudent)
	})

	return r
}
