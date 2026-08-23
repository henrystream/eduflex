package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type DrawdownService struct {
	repo *repository.DrawdownRepository
}

func NewDrawdownService(r *repository.DrawdownRepository) *DrawdownService {
	return &DrawdownService{repo: r}
}

type CreateDrawdownRequest struct {
	FacilityID pgtype.UUID    `json:"facility_id"`
	Amount     pgtype.Numeric `json:"amount"`
	Date       pgtype.Date    `json:"drawdown_date"`
	Reference  string         `json:"reference"`
}

func (s *DrawdownService) CreateDrawdown(ctx context.Context, req CreateDrawdownRequest) (db.LoanDrawdown, error) {
	if req.FacilityID.String() == "" {
		return db.LoanDrawdown{}, errors.New("facility_id required")
	}

	return s.repo.CreateDrawdown(ctx, repository.CreateDrawdownParams{
		FacilityID: req.FacilityID,
		Amount:     req.Amount,
		Date:       req.Date,
		Reference:  req.Reference,
	})
}

func (s *DrawdownService) ListDrawdownsByFacility(ctx context.Context, facilityID pgtype.UUID) ([]db.LoanDrawdown, error) {
	return s.repo.ListDrawdownsByFacility(ctx, facilityID)
}
