package http

import (
	"encoding/json"

	"net/http"

	"github.com/henrystream/eduflex/loan-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type FacilityHandler struct {
	svc *service.FacilityService
}

func NewFacilityHandler(svc *service.FacilityService) *FacilityHandler {
	return &FacilityHandler{svc: svc}
}

func (h *FacilityHandler) CreateFacility(w http.ResponseWriter, r *http.Request) {
	var req service.CreateFacilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	facility, err := h.svc.CreateFacility(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, facility)
}

func (h *FacilityHandler) ListFacilitiesByBank(w http.ResponseWriter, r *http.Request) {
	var buuid pgtype.UUID
	bankID := r.URL.Query().Get("bank_id")
	buuid.Scan(bankID)
	facilities, err := h.svc.ListFacilitiesByBank(r.Context(), buuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, facilities)
}
