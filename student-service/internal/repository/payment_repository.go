package repository

import (
	"context"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentRepository struct {
	queries *db.Queries
}

func NewPaymentRepository(q *db.Queries) *PaymentRepository {
	return &PaymentRepository{queries: q}
}

type CreatePaymentParams struct {
	InstallmentID        pgtype.UUID
	Amount               pgtype.Numeric
	PaymentMethod        pgtype.Text
	TransactionReference pgtype.Text
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, p CreatePaymentParams) (db.StudentPayment, error) {
	return r.queries.CreatePayment(ctx, db.CreatePaymentParams{
		InstallmentID:        p.InstallmentID,
		Amount:               p.Amount,
		PaymentMethod:        p.PaymentMethod,
		TransactionReference: p.TransactionReference,
	})
}

func (r *PaymentRepository) ListPaymentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentPayment, error) {
	return r.queries.ListPaymentsByStudent(ctx, studentID)
}
