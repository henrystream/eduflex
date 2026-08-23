package http

import (
	"encoding/json"

	"net/http"

	"github.com/henrystream/eduflex/loan-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepaymentHandler struct {
	svc *service.RepaymentService
}

func NewRepaymentHandler(svc *service.RepaymentService) *RepaymentHandler {
	return &RepaymentHandler{svc: svc}
}

func (h *RepaymentHandler) CreateRepayment(w http.ResponseWriter, r *http.Request) {
	var req service.CreateRepaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	repayment, err := h.svc.CreateRepayment(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, repayment)
}

func (h *RepaymentHandler) ListRepaymentsByDrawdown(w http.ResponseWriter, r *http.Request) {
	var duuid pgtype.UUID
	drawdownID := r.URL.Query().Get("drawdown_id")
	duuid.Scan(drawdownID)
	repayments, err := h.svc.ListRepaymentsByDrawdown(r.Context(), duuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, repayments)
}
