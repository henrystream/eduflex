package repository

import (
	"context"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type FacilityRepository struct {
	queries *db.Queries
}

func NewFacilityRepository(q *db.Queries) *FacilityRepository {
	return &FacilityRepository{queries: q}
}

type CreateFacilityParams struct {
	BankID       pgtype.UUID
	CreditLimit  pgtype.Numeric
	InterestRate pgtype.Numeric
	StartDate    pgtype.Date
	EndDate      pgtype.Date
	Status       string
}

func (r *FacilityRepository) CreateFacility(ctx context.Context, p CreateFacilityParams) (db.LoanFacility, error) {
	return r.queries.CreateFacility(ctx, db.CreateFacilityParams{
		BankID:       p.BankID,
		CreditLimit:  p.CreditLimit,
		InterestRate: p.InterestRate,
		StartDate:    p.StartDate,
		EndDate:      p.EndDate,
		Status:       p.Status,
	})
}

func (r *FacilityRepository) ListFacilitiesByBank(ctx context.Context, bankID pgtype.UUID) ([]db.LoanFacility, error) {
	return r.queries.ListFacilitiesByBank(ctx, bankID)
}
