package repository

import (
	"context"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type InstallmentRepository struct {
	queries *db.Queries
}

func NewInstallmentRepository(q *db.Queries) *InstallmentRepository {
	return &InstallmentRepository{queries: q}
}

type CreateInstallmentParams struct {
	FinancingID       pgtype.UUID
	InstallmentNumber int32
	DueDate           pgtype.Date
	Amount            pgtype.Numeric
	Status            string
}

func (r *InstallmentRepository) CreateInstallment(ctx context.Context, p CreateInstallmentParams) (db.MonthlyInstallment, error) {
	return r.queries.CreateInstallment(ctx, db.CreateInstallmentParams{
		FinancingID:       p.FinancingID,
		InstallmentNumber: p.InstallmentNumber,
		DueDate:           p.DueDate,
		Amount:            p.Amount,
		Status:            p.Status,
	})
}

func (r *InstallmentRepository) ListInstallments(ctx context.Context, financingID pgtype.UUID) ([]db.MonthlyInstallment, error) {
	return r.queries.ListInstallments(ctx, financingID)
}
