package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/api-gateway/internal/auth"
	"github.com/henrystream/eduflex/api-gateway/internal/config"
)

func NewRouter(cfg config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// JWT middleware (authenticates)
	jwt := auth.NewJWTMiddleware(cfg.JWTSecret)
	r.Use(jwt.Middleware)

	// PUBLIC ROUTES
	r.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		forward(w, r, cfg.AuthURL)
	})
	r.Post("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		forward(w, r, cfg.AuthURL)
	})

	// PROTECTED ROUTES WITH ROLE ENFORCEMENT

	// SCHOOL ROUTES
	r.Route("/schools", func(r chi.Router) {

		// Only admin can create schools
		r.With(auth.RequireRoles("admin")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})

		// Admin, school, finance can list schools
		r.With(auth.RequireRoles("admin", "school", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})

		// Admin, school, finance can view a school
		r.With(auth.RequireRoles("admin", "school", "finance")).Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.SchoolURL)
		})
	})

	// STUDENT ROUTES
	r.Route("/students", func(r chi.Router) {

		// Admin + school can create students
		r.With(auth.RequireRoles("admin", "school")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})

		// Admin, school, student can list students
		r.With(auth.RequireRoles("admin", "school", "student")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})

		// Admin, school, student can view a student
		r.With(auth.RequireRoles("admin", "school", "student")).Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
	})

	return r
}
