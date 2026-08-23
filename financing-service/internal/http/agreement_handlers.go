package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/financing-service/internal/service"
)

type AgreementHandler struct {
	svc *service.AgreementService
}

func NewAgreementHandler(svc *service.AgreementService) *AgreementHandler {
	return &AgreementHandler{svc: svc}
}

func (h *AgreementHandler) CreateAgreement(w http.ResponseWriter, r *http.Request) {
	var req service.CreateAgreementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	agreement, err := h.svc.CreateAgreement(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, agreement)
}
