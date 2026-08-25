package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/student-service/db/sqlc"
	"github.com/henrystream/eduflex/student-service/internal/events"
	"github.com/henrystream/eduflex/student-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentService struct {
	repo   *repository.PaymentRepository
	ledger LedgerClient
}

type LedgerClient interface {
	CreateEntry(req LedgerEntryRequest) error
}

type LedgerEntryRequest struct {
	EventType     string             `json:"event_type"`
	EventID       pgtype.UUID        `json:"event_id"`
	SourceService string             `json:"source_service"`
	DebitAccount  string             `json:"debit_account"`
	CreditAccount string             `json:"credit_account"`
	Amount        pgtype.Numeric     `json:"amount"`
	Currency      string             `json:"currency"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
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
	if req.InstallmentID.String() == "" || !req.Amount.Valid {
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
			SourceService: "student-service",
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

	var eventClient events.EventClient

	pubRequest := events.PublishEventRequest{
		EventType:     "STUDENT_PAYMENT_CREATED",
		SourceService: "student-service",
		AggregateID:   payment.ID,
		OccurredAt:    pgtype.Timestamptz{Time: payment.PaidAt.Time, Valid: true},
		Payload:       payment,
	}

	eventClient.Publish(pubRequest.EventType, pubRequest.AggregateID, pubRequest.OccurredAt, pubRequest.Payload)

	return payment, nil
}

func (s *PaymentService) ListPaymentsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.StudentPayment, error) {
	return s.repo.ListPaymentsByStudent(ctx, studentID)
}
