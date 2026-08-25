package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/events"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepaymentService struct {
	repo   *repository.RepaymentRepository
	ledger LedgerClient
}

func NewRepaymentService(r *repository.RepaymentRepository, ledger ...LedgerClient) *RepaymentService {
	var ledgerClient LedgerClient
	if len(ledger) > 0 {
		ledgerClient = ledger[0]
	}
	return &RepaymentService{repo: r, ledger: ledgerClient}
}

type CreateRepaymentRequest struct {
	DrawdownID pgtype.UUID    `json:"drawdown_id"`
	Amount     pgtype.Numeric `json:"amount"`
	Reference  string         `json:"reference"`
}

func (s *RepaymentService) CreateRepayment(ctx context.Context, req CreateRepaymentRequest) (db.LoanRepayment, error) {
	if req.DrawdownID.String() == "" {
		return db.LoanRepayment{}, errors.New("drawdown_id required")
	}

	repayment, err := s.repo.CreateRepayment(ctx, repository.CreateRepaymentParams{
		DrawdownID: req.DrawdownID,
		Amount:     req.Amount,
		Reference:  req.Reference,
	})
	if err != nil {
		return db.LoanRepayment{}, err
	}

	if s.ledger != nil {
		_ = s.ledger.CreateEntry(LedgerEntryRequest{
			EventType:     "LOAN_REPAYMENT",
			EventID:       repayment.ID,
			SourceService: "loan-service",
			DebitAccount:  "Cash - Bank",
			CreditAccount: "Loan Receivable - Bank",
			Amount:        repayment.Amount,
			Currency:      "AED",
			OccurredAt:    repayment.PaidAt,
		})
	}

	var eventClient events.EventClient

	pubRequest := events.PublishEventRequest{
		EventType:     "LOAN_REPAYMENT_CREATED",
		SourceService: "loan-service",
		AggregateID:   repayment.ID,
		OccurredAt:    pgtype.Timestamptz{Time: repayment.PaidAt.Time, Valid: true},
		Payload:       repayment,
	}

	eventClient.Publish(pubRequest.EventType, pubRequest.SourceService, pubRequest.AggregateID, pubRequest.OccurredAt, pubRequest.Payload)

	return repayment, nil
}

func (s *RepaymentService) ListRepaymentsByDrawdown(ctx context.Context, drawdownID pgtype.UUID) ([]db.LoanRepayment, error) {
	return s.repo.ListRepaymentsByDrawdown(ctx, drawdownID)
}
