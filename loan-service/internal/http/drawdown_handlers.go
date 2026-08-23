package http

import (
	"encoding/json"

	"net/http"

	"github.com/henrystream/eduflex/loan-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type DrawdownHandler struct {
	svc *service.DrawdownService
}

func NewDrawdownHandler(svc *service.DrawdownService) *DrawdownHandler {
	return &DrawdownHandler{svc: svc}
}

func (h *DrawdownHandler) CreateDrawdown(w http.ResponseWriter, r *http.Request) {
	var req service.CreateDrawdownRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	drawdown, err := h.svc.CreateDrawdown(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, drawdown)
}

func (h *DrawdownHandler) ListDrawdownsByFacility(w http.ResponseWriter, r *http.Request) {
	var fuuid pgtype.UUID
	facilityID := r.URL.Query().Get("facility_id")
	fuuid.Scan(facilityID)
	drawdowns, err := h.svc.ListDrawdownsByFacility(r.Context(), fuuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, drawdowns)
}
