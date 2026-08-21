package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/henrystream/eduflex/school-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type SchoolHandler struct {
	svc *service.SchoolService
}

// NewSchoolHandler creates a new instance of SchoolHandler.
func NewSchoolHandler(svc *service.SchoolService) *SchoolHandler {
	return &SchoolHandler{svc: svc}
}

func (h *SchoolHandler) CreateSchool(w http.ResponseWriter, r *http.Request) {

	var req service.CreateSchoolRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	school, err := h.svc.CreateSchool(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, school)

}

func (h *SchoolHandler) GetSchool(w http.ResponseWriter, r *http.Request) {
	var pgUUID pgtype.UUID
	id := chi.URLParam(r, "id")
	pgUUID.Scan(id)
	school, err := h.svc.GetSchool(context.Background(), pgUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeJSON(w, http.StatusOK, school)

}

func (h *SchoolHandler) ListSchools(w http.ResponseWriter, r *http.Request) {
	schools, err := h.svc.ListSchools(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, schools)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
