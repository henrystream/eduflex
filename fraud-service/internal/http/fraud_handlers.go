package http

import (
	"encoding/json"
	"net/http"

	"github.com/henrystream/eduflex/fraud-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type FraudHandler struct {
	svc *service.FraudService
}

func NewFraudHandler(svc *service.FraudService) *FraudHandler {
	return &FraudHandler{svc: svc}
}

func (h *FraudHandler) CheckAgreement(w http.ResponseWriter, r *http.Request) {
	var agreeUUID, studentUUID pgtype.UUID
	agreementID := r.URL.Query().Get("agreement_id")
	studentID := r.URL.Query().Get("student_id")
	err := studentUUID.Scan(studentID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = agreeUUID.Scan(agreementID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.svc.CheckAgreementRisk(studentUUID, agreeUUID)
	if err != nil {
		//http.Error(w, "Error is here", http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
