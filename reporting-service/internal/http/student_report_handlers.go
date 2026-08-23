package http

import (
	"net/http"

	"github.com/henrystream/eduflex/reporting-service/internal/service"
)

type StudentReportHandler struct {
	svc *service.StudentReportService
}

func NewStudentReportHandler(svc *service.StudentReportService) *StudentReportHandler {
	return &StudentReportHandler{svc: svc}
}

func (h *StudentReportHandler) GetStudentLoanSummary(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_id")
	if studentID == "" {
		http.Error(w, "student_id query parameter is required", http.StatusBadRequest)
		return
	}

	report, err := h.svc.GenerateStudentLoanSummary(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, report)
}
