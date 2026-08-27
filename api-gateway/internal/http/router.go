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

	// ENROLLMENT ROUTES
	r.Route("/enrollments", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "school", "student")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
		r.With(auth.RequireRoles("admin", "school", "student")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
	})

	// PAYMENT ROUTES
	r.Route("/payments", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "school", "student")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
		r.With(auth.RequireRoles("admin", "school", "student")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.StudentURL)
		})
	})

	// FINANCING ROUTES
	r.Route("/agreements", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.FinancingURL)
		})
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.FinancingURL)
		})
	})

	r.Route("/installments", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.FinancingURL)
		})
	})

	// LOAN ROUTES
	r.Route("/banks", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
	})

	r.Route("/facilities", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
	})

	r.Route("/drawdowns", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
	})

	r.Route("/repayments", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
		r.With(auth.RequireRoles("admin", "finance")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.LoanURL)
		})
	})

	// DISBURSEMENT ROUTES
	r.Route("/disbursements", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance", "school")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.DisbursementURL)
		})
		r.With(auth.RequireRoles("admin", "finance", "school")).Get("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.DisbursementURL)
		})
	})

	// REPORTING ROUTES
	r.Route("/reports", func(r chi.Router) {
		// Admin, finance, school can access school reports
		r.With(auth.RequireRoles("admin", "finance", "school")).Get("/school", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.ReportingURL)
		})

		// Admin, finance, student can access student reports
		r.With(auth.RequireRoles("admin", "finance", "student")).Get("/student", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.ReportingURL)
		})

		// Only admin and finance can access financial reports
		r.With(auth.RequireRoles("admin", "finance")).Get("/financial", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.ReportingURL)
		})
	})

	//EVENTS ROUTES

	r.Route("/events", func(r chi.Router) {
		r.With(auth.RequireRoles("admin", "finance", "school")).Post("/", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.EventsURL)
		})
		r.With(auth.RequireRoles("admin", "finance", "school")).Get("/unprocessed", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.EventsURL)
		})
	})

	//FRAUD ROUTES
	r.Route("/fraud", func(r chi.Router) {

		r.With(auth.RequireRoles("admin", "finance", "school")).Get("/check-agreement", func(w http.ResponseWriter, r *http.Request) {
			forward(w, r, cfg.FraudURL)
		})
	})

	return r
}
