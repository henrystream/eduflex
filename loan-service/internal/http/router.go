package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/loan-service/internal/service"
)

func NewRouter(
	bankSvc *service.BankService,
	facilitySvc *service.FacilityService,
	drawdownSvc *service.DrawdownService,
	repaymentSvc *service.RepaymentService,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	bankHandler := NewBankHandler(bankSvc)
	facilityHandler := NewFacilityHandler(facilitySvc)
	drawdownHandler := NewDrawdownHandler(drawdownSvc)
	repaymentHandler := NewRepaymentHandler(repaymentSvc)

	r.Route("/banks", func(r chi.Router) {
		r.Post("/", bankHandler.CreateBank)
		r.Get("/", bankHandler.ListBanks)
	})

	r.Route("/facilities", func(r chi.Router) {
		r.Post("/", facilityHandler.CreateFacility)
		r.Get("/", facilityHandler.ListFacilitiesByBank)
	})

	r.Route("/drawdowns", func(r chi.Router) {
		r.Post("/", drawdownHandler.CreateDrawdown)
		r.Get("/", drawdownHandler.ListDrawdownsByFacility)
	})

	r.Route("/repayments", func(r chi.Router) {
		r.Post("/", repaymentHandler.CreateRepayment)
		r.Get("/", repaymentHandler.ListRepaymentsByDrawdown)
	})

	return r
}
