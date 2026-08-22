package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/student-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type EnrollmentHandler struct {
	svc *service.EnrollmentService
}

func NewEnrollmentHandler(svc *service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{svc: svc}
}

func (h *EnrollmentHandler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var req service.CreateEnrollmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	enrollment, err := h.svc.CreateEnrollment(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, enrollment)
}

func (h *EnrollmentHandler) ListEnrollmentsByStudent(w http.ResponseWriter, r *http.Request) {
	var suuid pgtype.UUID
	studentID := r.URL.Query().Get("student_id")
	suuid.Scan(studentID)

	enrollments, err := h.svc.ListEnrollmentsByStudent(r.Context(), suuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, enrollments)
}
