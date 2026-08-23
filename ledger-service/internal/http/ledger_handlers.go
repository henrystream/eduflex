package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/ledger-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type LedgerHandler struct {
	svc *service.LedgerService
}

func NewLedgerHandler(svc *service.LedgerService) *LedgerHandler {
	return &LedgerHandler{svc: svc}
}

func (h *LedgerHandler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var req service.CreateLedgerEntryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	entry, err := h.svc.CreateEntry(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, entry)
}

func (h *LedgerHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	entries, err := h.svc.ListAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

func (h *LedgerHandler) ListByEvent(w http.ResponseWriter, r *http.Request) {
	eventType := r.URL.Query().Get("event_type")
	eventID := r.URL.Query().Get("event_id")
	var euuid pgtype.UUID
	euuid.Scan(eventID)
	entries, err := h.svc.ListByEvent(r.Context(), eventType, euuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

func (h *LedgerHandler) ListByService(w http.ResponseWriter, r *http.Request) {
	sourceService := r.URL.Query().Get("source_service")

	entries, err := h.svc.ListByService(r.Context(), sourceService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
