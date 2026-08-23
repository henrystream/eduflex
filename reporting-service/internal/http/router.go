package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/reporting-service/internal/service"
)

func NewRouter(
	schoolSvc *service.SchoolReportService,
	studentSvc *service.StudentReportService,
	financialSvc *service.FinancialReportService,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	schoolHandler := NewSchoolReportHandler(schoolSvc)
	studentHandler := NewStudentReportHandler(studentSvc)
	financialHandler := NewFinancialReportHandler(financialSvc)

	r.Get("/reports/school", schoolHandler.GetSchoolStatement)
	r.Get("/reports/student", studentHandler.GetStudentLoanSummary)
	r.Get("/reports/financial", financialHandler.GetFinancialStatement)

	return r
}
