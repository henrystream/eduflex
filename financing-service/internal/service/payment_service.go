package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/events"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentService struct {
	repo   *repository.PaymentRepository
	ledger LedgerClient
}

func NewPaymentService(r *repository.PaymentRepository, ledger ...LedgerClient) *PaymentService {
	var ledgerClient LedgerClient
	if len(ledger) > 0 {
		ledgerClient = ledger[0]
	}
	return &PaymentService{repo: r, ledger: ledgerClient}
}

type CreatePaymentRequest struct {
	InstallmentID        pgtype.UUID    `json:"installment_id"`
	Amount               pgtype.Numeric `json:"amount"`
	PaymentMethod        pgtype.Text    `json:"payment_method"`
	TransactionReference pgtype.Text    `json:"transaction_reference"`
}

func (s *PaymentService) CreatePayment(ctx context.Context, req CreatePaymentRequest) (db.StudentPayment, error) {
	if !req.InstallmentID.Valid || !req.Amount.Valid {
		return db.StudentPayment{}, errors.New("installment_id and amount required")
	}

	payment, err := s.repo.CreatePayment(ctx, repository.CreatePaymentParams{
		InstallmentID:        req.InstallmentID,
		Amount:               req.Amount,
		PaymentMethod:        req.PaymentMethod,
		TransactionReference: req.TransactionReference,
	})
	if err != nil {
		return db.StudentPayment{}, err
	}

	if s.ledger != nil {
		_ = s.ledger.CreateEntry(LedgerEntryRequest{
			EventType:     "STUDENT_PAYMENT",
			EventID:       payment.ID,
			SourceService: "financing-service",
			DebitAccount:  "Cash - Bank",
			CreditAccount: "Accounts Receivable - Student",
			Amount:        payment.Amount,
			Currency:      "AED",
			OccurredAt: pgtype.Timestamptz{
				Time:  payment.PaidAt.Time,
				Valid: payment.PaidAt.Valid,
			},
		})
	}

	events.NewEventClient("http://event-service:8080", "financing-service").Publish(
		"STUDENT_PAYMENT_CREATED",
		"financing-service",
		payment.ID,
		pgtype.Timestamptz{Time: payment.PaidAt.Time, Valid: payment.PaidAt.Valid},
		payment,
	)

	return payment, nil
}

func (s *PaymentService) ListPaymentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentPayment, error) {
	return s.repo.ListPaymentsByStudent(ctx, studentID)
}
