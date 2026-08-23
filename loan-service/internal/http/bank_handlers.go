package http

import (
	"encoding/json"

	"net/http"

	"github.com/henrystream/eduflex/loan-service/internal/service"
)

type BankHandler struct {
	svc *service.BankService
}

func NewBankHandler(svc *service.BankService) *BankHandler {
	return &BankHandler{svc: svc}
}

func (h *BankHandler) CreateBank(w http.ResponseWriter, r *http.Request) {
	var req service.CreateBankRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	bank, err := h.svc.CreateBank(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, bank)
}

func (h *BankHandler) ListBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := h.svc.ListBanks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, banks)
}
