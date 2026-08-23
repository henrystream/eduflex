package repository

import (
	"context"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AgreementRepository struct {
	queries *db.Queries
}

func NewAgreementRepository(q *db.Queries) *AgreementRepository {
	return &AgreementRepository{queries: q}
}

type CreateAgreementParams struct {
	StudentID    pgtype.UUID
	InvoiceID    pgtype.UUID
	Principal    pgtype.Numeric
	InterestRate pgtype.Numeric
	ServiceFee   pgtype.Numeric
	TotalPayable pgtype.Numeric
	TermMonths   int32
	StartDate    pgtype.Date
	Status       string
}

func (r *AgreementRepository) CreateAgreement(ctx context.Context, p CreateAgreementParams) (db.FinancingAgreement, error) {
	return r.queries.CreateAgreement(ctx, db.CreateAgreementParams{
		StudentID:    p.StudentID,
		InvoiceID:    p.InvoiceID,
		Principal:    p.Principal,
		InterestRate: p.InterestRate,
		ServiceFee:   p.ServiceFee,
		TotalPayable: p.TotalPayable,
		TermMonths:   p.TermMonths,
		StartDate:    p.StartDate,
		Status:       p.Status,
	})
}

func (r *AgreementRepository) GetAgreement(ctx context.Context, id pgtype.UUID) (db.FinancingAgreement, error) {
	return r.queries.GetAgreement(ctx, id)
}

func (r *AgreementRepository) ListAgreementsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.FinancingAgreement, error) {
	return r.queries.ListAgreementsByStudent(ctx, studentID)
}
