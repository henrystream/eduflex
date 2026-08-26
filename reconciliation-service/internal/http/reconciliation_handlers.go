package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/reconciliation-service/internal/service"
)

type ReconciliationHandler struct {
	svc *service.ReconciliationService
}

func NewReconciliationHandler(svc *service.ReconciliationService) *ReconciliationHandler {
	return &ReconciliationHandler{svc: svc}
}

func (h *ReconciliationHandler) Reconcile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("file")

	results, err := h.svc.Reconcile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, results)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
