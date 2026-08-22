package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type EnrollmentService struct {
	repo *repository.EnrollmentRepository
}

func NewEnrollmentService(r *repository.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{repo: r}
}

type CreateEnrollmentRequest struct {
	StudentID      pgtype.UUID `json:"student_id"`
	SchoolID       pgtype.UUID `json:"school_id"`
	EnrollmentDate pgtype.Date `json:"enrollment_date"`
	Status         string      `json:"status"`
}

func (s *EnrollmentService) CreateEnrollment(ctx context.Context, req CreateEnrollmentRequest) (db.StudentSchoolEnrollment, error) {
	if req.StudentID.String() == "" || req.SchoolID.String() == "" {
		return db.StudentSchoolEnrollment{}, errors.New("student_id and school_id required")
	}

	return s.repo.CreateEnrollment(ctx, repository.CreateEnrollmentParams{
		StudentID:      req.StudentID,
		SchoolID:       req.SchoolID,
		EnrollmentDate: req.EnrollmentDate,
		Status:         req.Status,
	})
}

func (s *EnrollmentService) ListEnrollmentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentSchoolEnrollment, error) {
	return s.repo.ListEnrollmentsByStudent(ctx, studentID)
}
