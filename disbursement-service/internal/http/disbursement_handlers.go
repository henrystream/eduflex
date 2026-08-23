package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/disbursement-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type DisbursementHandler struct {
	svc *service.DisbursementService
}

func NewDisbursementHandler(svc *service.DisbursementService) *DisbursementHandler {
	return &DisbursementHandler{svc: svc}
}

func (h *DisbursementHandler) CreateDisbursement(w http.ResponseWriter, r *http.Request) {
	var req service.CreateDisbursementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	disbursement, err := h.svc.CreateDisbursement(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, disbursement)
}

func (h *DisbursementHandler) ListDisbursementsBySchool(w http.ResponseWriter, r *http.Request) {
	var schoolUUID pgtype.UUID
	schoolID := r.URL.Query().Get("school_id")
	if err := schoolUUID.Scan(schoolID); err != nil {
		http.Error(w, "invalid school_id", http.StatusBadRequest)
		return
	}

	disbursements, err := h.svc.ListDisbursementsBySchool(r.Context(), schoolUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, disbursements)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
