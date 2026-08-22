package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/henrystream/eduflex/auth-service/internal/service"
)

func NewRouter(svc *service.AuthService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewAuthHandler(svc)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	return r
}
