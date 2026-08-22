package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/student-service/internal/service"
)

func NewRouter(svc *service.StudentService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewStudentHandler(svc)

	r.Route("/students", func(r chi.Router) {
		r.Post("/", h.CreateStudent)
		r.Get("/", h.ListStudents)
		r.Get("/{id}", h.GetStudent)
	})

	return r
}
