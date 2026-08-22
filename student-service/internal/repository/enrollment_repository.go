package repository

import (
	"context"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type EnrollmentRepository struct {
	queries *db.Queries
}

func NewEnrollmentRepository(q *db.Queries) *EnrollmentRepository {
	return &EnrollmentRepository{queries: q}
}

type CreateEnrollmentParams struct {
	StudentID      pgtype.UUID
	SchoolID       pgtype.UUID
	EnrollmentDate pgtype.Date
	Status         string
}

func (r *EnrollmentRepository) CreateEnrollment(ctx context.Context, p CreateEnrollmentParams) (db.StudentSchoolEnrollment, error) {
	return r.queries.CreateEnrollment(ctx, db.CreateEnrollmentParams{
		StudentID:      p.StudentID,
		SchoolID:       p.SchoolID,
		EnrollmentDate: p.EnrollmentDate,
		Status:         p.Status,
	})
}

func (r *EnrollmentRepository) ListEnrollmentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentSchoolEnrollment, error) {
	return r.queries.ListEnrollmentsByStudent(ctx, studentID)
}
