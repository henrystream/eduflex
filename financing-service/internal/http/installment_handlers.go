package http

import (
	"net/http"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type InstallmentHandler struct {
	svc *service.InstallmentService
}

func NewInstallmentHandler(svc *service.InstallmentService) *InstallmentHandler {
	return &InstallmentHandler{svc: svc}
}

func (h *InstallmentHandler) ListInstallments(w http.ResponseWriter, r *http.Request) {
	financingID := r.URL.Query().Get("financing_id")
	if financingID == "" {
		http.Error(w, "financing_id is required", http.StatusBadRequest)
		return
	}

	var fid pgtype.UUID
	if err := fid.Scan(financingID); err != nil {
		http.Error(w, "invalid financing_id", http.StatusBadRequest)
		return
	}

	installments, err := h.svc.ListInstallments(r.Context(), fid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if installments == nil {
		installments = []db.MonthlyInstallment{}
	}

	writeJSON(w, http.StatusOK, installments)
}
