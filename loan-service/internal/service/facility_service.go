package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type FacilityService struct {
	repo *repository.FacilityRepository
}

func NewFacilityService(r *repository.FacilityRepository) *FacilityService {
	return &FacilityService{repo: r}
}

type CreateFacilityRequest struct {
	BankID       pgtype.UUID    `json:"bank_id"`
	CreditLimit  pgtype.Numeric `json:"credit_limit"`
	InterestRate pgtype.Numeric `json:"interest_rate"`
	StartDate    pgtype.Date    `json:"start_date"`
	EndDate      pgtype.Date    `json:"end_date"`
}

func (s *FacilityService) CreateFacility(ctx context.Context, req CreateFacilityRequest) (db.LoanFacility, error) {
	if req.BankID.String() == "" {
		return db.LoanFacility{}, errors.New("bank_id is required")
	}

	return s.repo.CreateFacility(ctx, repository.CreateFacilityParams{
		BankID:       req.BankID,
		CreditLimit:  req.CreditLimit,
		InterestRate: req.InterestRate,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Status:       "ACTIVE",
	})
}

func (s *FacilityService) ListFacilitiesByBank(ctx context.Context, bankID pgtype.UUID) ([]db.LoanFacility, error) {
	return s.repo.ListFacilitiesByBank(ctx, bankID)
}
