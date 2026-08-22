package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/henrystream/eduflex/student-service/internal/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type StudentHandler struct {
	svc *service.StudentService
}

func NewStudentHandler(svc *service.StudentService) *StudentHandler {
	return &StudentHandler{svc: svc}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var req service.CreateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	student, err := h.svc.CreateStudent(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, student)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	var pgUUID pgtype.UUID
	id := chi.URLParam(r, "id")
	pgUUID.Scan(id)

	student, err := h.svc.GetStudent(r.Context(), pgUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, student)
}

func (h *StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.svc.ListStudents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, students)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
