package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/auth-service/db/sqlc"
	"github.com/henrystream/eduflex/auth-service/internal/repository"
	"github.com/henrystream/eduflex/auth-service/internal/token"
)

type AuthService struct {
	repo     *repository.UserRepository
	jwtMaker *token.JWTMaker
}

func NewAuthService(repo *repository.UserRepository, jwtMaker *token.JWTMaker) *AuthService {
	return &AuthService{repo: repo, jwtMaker: jwtMaker}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (db.User, error) {
	if req.Email == "" || req.Password == "" {
		return db.User{}, errors.New("email and password required")
	}

	hash, err := token.HashPassword(req.Password)
	if err != nil {
		return db.User{}, err
	}

	return s.repo.CreateUser(ctx, req.Email, hash, req.Role)
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := token.CheckPasswordHash(req.Password, user.PasswordHash); err != nil {
		return "", errors.New("invalid email or password")
	}

	return s.jwtMaker.GenerateToken(user.ID.String(), user.Role)
}
