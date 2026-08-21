package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/school-service/db/sqlc"
	"github.com/henrystream/eduflex/school-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type SchoolService struct {
	repo *repository.SchoolRepository
}

func NewSchoolService(r *repository.SchoolRepository) *SchoolService {
	return &SchoolService{repo: r}
}

type CreateSchoolRequest struct {
	Name         string      `json:"name"`
	Address      pgtype.Text `json:"address"`
	ContactEmail pgtype.Text `json:"contact_email"`
	ContactPhone pgtype.Text `json:"contact_phone"`
}

func (s *SchoolService) CreateSchool(ctx context.Context, req CreateSchoolRequest) (db.School, error) {
	if req.Name == "" {
		return db.School{}, errors.New("name is required")
	}

	return s.repo.CreateSchool(ctx, repository.CreateSchoolParams{
		Name:         req.Name,
		Address:      req.Address,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
	})
}

func (s *SchoolService) GetSchool(ctx context.Context, id pgtype.UUID) (db.School, error) {
	return s.repo.GetSchool(ctx, id)
}

func (s *SchoolService) ListSchools(ctx context.Context) ([]db.School, error) {
	return s.repo.ListSchools(ctx)
}
