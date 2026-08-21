package http

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/henrystream/eduflex/school-service/internal/service"
)

func NewRouter(svc *service.SchoolService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewSchoolHandler(svc)

	r.Route("/schools", func(r chi.Router) {
		r.Post("/", h.CreateSchool)
		r.Get("/", h.ListSchools)
		r.Get("/{id}", h.GetSchool)
	})

	return r
}
