package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/student-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentHandler struct {
	svc *service.PaymentService
}

func NewPaymentHandler(svc *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{svc: svc}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req service.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	payment, err := h.svc.CreatePayment(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, payment)
}

func (h *PaymentHandler) ListPaymentsByStudent(w http.ResponseWriter, r *http.Request) {
	var suuid pgtype.UUID
	studentID := r.URL.Query().Get("student_id")
	suuid.Scan(studentID)

	payments, err := h.svc.ListPaymentsByStudent(r.Context(), suuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, payments)
}
