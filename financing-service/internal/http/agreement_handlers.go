package http

import (
	"encoding/json"
	"net/http"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
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

func (h *AgreementHandler) ListAgreementsByStudent(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_id")
	if studentID == "" {
		http.Error(w, "student_id is required", http.StatusBadRequest)
		return
	}

	var sid pgtype.UUID
	if err := sid.Scan(studentID); err != nil {
		http.Error(w, "invalid student_id", http.StatusBadRequest)
		return
	}

	agreements, err := h.svc.ListAgreementsByStudent(r.Context(), sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if agreements == nil {
		agreements = []db.FinancingAgreement{}
	}

	writeJSON(w, http.StatusOK, agreements)
}
