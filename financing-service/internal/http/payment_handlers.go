package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/financing-service/internal/service"
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
	var studentID pgtype.UUID
	if err := studentID.Scan(r.URL.Query().Get("student_id")); err != nil {
		http.Error(w, "invalid student_id", http.StatusBadRequest)
		return
	}

	payments, err := h.svc.ListPaymentsByStudent(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, payments)
}
