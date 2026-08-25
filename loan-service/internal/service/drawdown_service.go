package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/events"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type DrawdownService struct {
	repo   *repository.DrawdownRepository
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

func NewDrawdownService(r *repository.DrawdownRepository, ledger ...LedgerClient) *DrawdownService {
	var ledgerClient LedgerClient
	if len(ledger) > 0 {
		ledgerClient = ledger[0]
	}
	return &DrawdownService{repo: r, ledger: ledgerClient}
}

type CreateDrawdownRequest struct {
	FacilityID pgtype.UUID    `json:"facility_id"`
	Amount     pgtype.Numeric `json:"amount"`
	Date       pgtype.Date    `json:"drawdown_date"`
	Reference  string         `json:"reference"`
}

func (s *DrawdownService) CreateDrawdown(ctx context.Context, req CreateDrawdownRequest) (db.LoanDrawdown, error) {
	if req.FacilityID.String() == "" {
		return db.LoanDrawdown{}, errors.New("facility_id required")
	}

	drawdown, err := s.repo.CreateDrawdown(ctx, repository.CreateDrawdownParams{
		FacilityID: req.FacilityID,
		Amount:     req.Amount,
		Date:       req.Date,
		Reference:  req.Reference,
	})
	if err != nil {
		return db.LoanDrawdown{}, err
	}

	if s.ledger != nil {
		_ = s.ledger.CreateEntry(LedgerEntryRequest{
			EventType:     "LOAN_DRAWDOWN",
			EventID:       drawdown.ID,
			SourceService: "loan-service",
			DebitAccount:  "Loan Receivable - Bank",
			CreditAccount: "Cash - Bank Facility",
			Amount:        drawdown.Amount,
			Currency:      "AED",
			OccurredAt: pgtype.Timestamptz{
				Time:  drawdown.DrawdownDate.Time,
				Valid: drawdown.DrawdownDate.Valid,
			},
		})
	}

	var eventClient events.EventClient

	pubRequest := events.PublishEventRequest{
		EventType:     "LOAN_DRAWDOWN_CREATED",
		SourceService: "loan-service",
		AggregateID:   drawdown.ID,
		OccurredAt:    pgtype.Timestamptz{Time: drawdown.DrawdownDate.Time, Valid: true},
		Payload:       drawdown,
	}

	eventClient.Publish(pubRequest.EventType, pubRequest.AggregateID, pubRequest.OccurredAt, pubRequest.Payload)

	return drawdown, nil
}

func (s *DrawdownService) ListDrawdownsByFacility(ctx context.Context, facilityID pgtype.UUID) ([]db.LoanDrawdown, error) {
	return s.repo.ListDrawdownsByFacility(ctx, facilityID)
}
