package http

import (
	"net/http"

	"github.com/henrystream/eduflex/reporting-service/internal/service"
)

type FinancialReportHandler struct {
	svc *service.FinancialReportService
}

func NewFinancialReportHandler(svc *service.FinancialReportService) *FinancialReportHandler {
	return &FinancialReportHandler{svc: svc}
}

func (h *FinancialReportHandler) GetFinancialStatement(w http.ResponseWriter, r *http.Request) {
	report, err := h.svc.GenerateFinancialStatement()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, report)
}
