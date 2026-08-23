package http

import (
	"net/http"

	"github.com/henrystream/eduflex/reporting-service/internal/service"
)

type SchoolReportHandler struct {
	svc *service.SchoolReportService
}

func NewSchoolReportHandler(svc *service.SchoolReportService) *SchoolReportHandler {
	return &SchoolReportHandler{svc: svc}
}

func (h *SchoolReportHandler) GetSchoolStatement(w http.ResponseWriter, r *http.Request) {
	schoolID := r.URL.Query().Get("school_id")
	if schoolID == "" {
		http.Error(w, "school_id query parameter is required", http.StatusBadRequest)
		return
	}

	report, err := h.svc.GenerateSchoolStatement(schoolID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, report)
}
