package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(r *repository.StudentRepository) *StudentService {
	return &StudentService{repo: r}
}

type CreateStudentRequest struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	DateOfBirth pgtype.Date `json:"date_of_birth"`
	Email       string      `json:"email"`
	Phone       pgtype.Text `json:"phone"`
}

func (s *StudentService) CreateStudent(ctx context.Context, req CreateStudentRequest) (db.Student, error) {
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		return db.Student{}, errors.New("first_name, last_name and email are required")
	}

	return s.repo.CreateStudent(ctx, repository.CreateStudentParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Email:       req.Email,
		Phone:       req.Phone,
	})
}

func (s *StudentService) GetStudent(ctx context.Context, id pgtype.UUID) (db.Student, error) {
	return s.repo.GetStudent(ctx, id)
}

func (s *StudentService) ListStudents(ctx context.Context) ([]db.Student, error) {
	return s.repo.ListStudents(ctx)
}
