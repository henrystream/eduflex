package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(r *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: r}
}

type CreatePaymentRequest struct {
	InstallmentID        pgtype.UUID    `json:"installment_id"`
	Amount               pgtype.Numeric `json:"amount"`
	PaymentMethod        pgtype.Text    `json:"payment_method"`
	TransactionReference pgtype.Text    `json:"transaction_reference"`
}

func (s *PaymentService) CreatePayment(ctx context.Context, req CreatePaymentRequest) (db.StudentPayment, error) {
	if req.InstallmentID.String() == "" || req.Amount.InfinityModifier.String() == "" {
		return db.StudentPayment{}, errors.New("installment_id and amount required")
	}

	return s.repo.CreatePayment(ctx, repository.CreatePaymentParams{
		InstallmentID:        req.InstallmentID,
		Amount:               req.Amount,
		PaymentMethod:        req.PaymentMethod,
		TransactionReference: req.TransactionReference,
	})
}

func (s *PaymentService) ListPaymentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentPayment, error) {
	return s.repo.ListPaymentsByStudent(ctx, studentID)
}
