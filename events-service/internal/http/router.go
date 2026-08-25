package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/events-service/internal/service"
)

func NewRouter(svc *service.EventService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewEventHandler(svc)

	r.Route("/events", func(r chi.Router) {
		r.Post("/", h.Publish)
		r.Get("/unprocessed", h.ListUnprocessed)
	})

	return r
}
