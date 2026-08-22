package repository

import (
	"context"

	db "github.com/henrystream/eduflex/auth-service/db/sqlc"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{queries: q}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, passwordHash, role string) (db.User, error) {
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	})
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}
