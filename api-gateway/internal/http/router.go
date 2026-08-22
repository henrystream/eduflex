package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/api-gateway/internal/config"
)

func NewRouter(cfg config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Auth routes
	r.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		forward(w, r, cfg.AuthURL)
	})
	r.Post("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		forward(w, r, cfg.AuthURL)
	})

	// School routes
	r.Route("/schools", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})
	})

	// Student routes
	r.Route("/students", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
	})

	return r
}
