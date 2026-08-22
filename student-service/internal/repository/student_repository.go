package repository

import (
	"context"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type StudentRepository struct {
	queries *db.Queries
}

func NewStudentRepository(q *db.Queries) *StudentRepository {
	return &StudentRepository{queries: q}
}

type CreateStudentParams struct {
	FirstName   string
	LastName    string
	DateOfBirth pgtype.Date // could be time.Time if you prefer
	Email       string
	Phone       pgtype.Text
}

func (r *StudentRepository) CreateStudent(ctx context.Context, p CreateStudentParams) (db.Student, error) {
	return r.queries.CreateStudent(ctx, db.CreateStudentParams{
		FirstName:   p.FirstName,
		LastName:    p.LastName,
		DateOfBirth: p.DateOfBirth,
		Email:       p.Email,
		Phone:       p.Phone,
	})
}

func (r *StudentRepository) GetStudent(ctx context.Context, id pgtype.UUID) (db.Student, error) {
	return r.queries.GetStudent(ctx, id)
}

func (r *StudentRepository) ListStudents(ctx context.Context) ([]db.Student, error) {
	return r.queries.ListStudents(ctx)
}
