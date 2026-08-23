package repository

import (
	"context"

	db "github.com/henrystream/eduflex/disbursement-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type DisbursementRepository struct {
	queries *db.Queries
}

func NewDisbursementRepository(q *db.Queries) *DisbursementRepository {
	return &DisbursementRepository{queries: q}
}

type CreateDisbursementParams struct {
	SchoolID      pgtype.UUID
	InvoiceID     pgtype.UUID
	Amount        pgtype.Numeric
	PaymentMethod string
	Reference     string
	Status        string
}

func (r *DisbursementRepository) CreateDisbursement(ctx context.Context, p CreateDisbursementParams) (db.EduflexDisbursement, error) {
	return r.queries.CreateDisbursement(ctx, db.CreateDisbursementParams{
		SchoolID:      p.SchoolID,
		InvoiceID:     p.InvoiceID,
		Amount:        p.Amount,
		PaymentMethod: p.PaymentMethod,
		Reference:     p.Reference,
		Status:        p.Status,
	})
}

func (r *DisbursementRepository) ListDisbursementsBySchool(ctx context.Context, schoolID pgtype.UUID) ([]db.EduflexDisbursement, error) {
	return r.queries.ListDisbursementsBySchool(ctx, schoolID)
}
