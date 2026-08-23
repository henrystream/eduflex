package repository

import (
	"context"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepaymentRepository struct {
	queries *db.Queries
}

func NewRepaymentRepository(q *db.Queries) *RepaymentRepository {
	return &RepaymentRepository{queries: q}
}

type CreateRepaymentParams struct {
	DrawdownID pgtype.UUID
	Amount     pgtype.Numeric
	Reference  string
}

func (r *RepaymentRepository) CreateRepayment(ctx context.Context, p CreateRepaymentParams) (db.LoanRepayment, error) {
	return r.queries.CreateRepayment(ctx, db.CreateRepaymentParams{
		DrawdownID: p.DrawdownID,
		Amount:     p.Amount,
		Reference:  p.Reference,
	})
}

func (r *RepaymentRepository) ListRepaymentsByDrawdown(ctx context.Context, drawdownID pgtype.UUID) ([]db.LoanRepayment, error) {
	return r.queries.ListRepaymentsByDrawdown(ctx, drawdownID)
}
