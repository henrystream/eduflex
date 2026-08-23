package repository

import (
	"context"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type DrawdownRepository struct {
	queries *db.Queries
}

func NewDrawdownRepository(q *db.Queries) *DrawdownRepository {
	return &DrawdownRepository{queries: q}
}

type CreateDrawdownParams struct {
	FacilityID pgtype.UUID
	Amount     pgtype.Numeric
	Date       pgtype.Date
	Reference  string
}

func (r *DrawdownRepository) CreateDrawdown(ctx context.Context, p CreateDrawdownParams) (db.LoanDrawdown, error) {
	return r.queries.CreateDrawdown(ctx, db.CreateDrawdownParams{
		FacilityID:   p.FacilityID,
		Amount:       p.Amount,
		DrawdownDate: p.Date,
		Reference:    p.Reference,
	})
}

func (r *DrawdownRepository) ListDrawdownsByFacility(ctx context.Context, facilityID pgtype.UUID) ([]db.LoanDrawdown, error) {
	return r.queries.ListDrawdownsByFacility(ctx, facilityID)
}
