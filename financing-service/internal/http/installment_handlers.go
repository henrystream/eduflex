package http

import (
	"net/http"

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
	var fid pgtype.UUID
	fid.Scan(financingID)
	installments, err := h.svc.ListInstallments(r.Context(), fid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, installments)
}
