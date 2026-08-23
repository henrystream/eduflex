package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/ledger-service/db/sqlc"
	"github.com/henrystream/eduflex/ledger-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type LedgerService struct {
	repo *repository.LedgerRepository
}

func NewLedgerService(r *repository.LedgerRepository) *LedgerService {
	return &LedgerService{repo: r}
}

type CreateLedgerEntryRequest struct {
	EventType     string             `json:"event_type"`
	EventID       pgtype.UUID        `json:"event_id"`
	SourceService string             `json:"source_service"`
	DebitAccount  string             `json:"debit_account"`
	CreditAccount string             `json:"credit_account"`
	Amount        pgtype.Numeric     `json:"amount"`
	Currency      string             `json:"currency"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}

func (s *LedgerService) CreateEntry(ctx context.Context, req CreateLedgerEntryRequest) (db.LedgerEntry, error) {
	if req.EventType == "" || req.EventID.String() == "" {
		return db.LedgerEntry{}, errors.New("event_type and event_id required")
	}

	return s.repo.CreateLedgerEntry(ctx, repository.CreateLedgerEntryParams{
		EventType:     req.EventType,
		EventID:       req.EventID,
		SourceService: req.SourceService,
		DebitAccount:  req.DebitAccount,
		CreditAccount: req.CreditAccount,
		Amount:        req.Amount,
		Currency:      req.Currency,
		OccurredAt:    req.OccurredAt,
	})
}

func (s *LedgerService) ListAll(ctx context.Context) ([]db.LedgerEntry, error) {
	return s.repo.ListLedgerEntries(ctx)
}

func (s *LedgerService) ListByEvent(ctx context.Context, eventType string, eventID pgtype.UUID) ([]db.LedgerEntry, error) {
	return s.repo.ListLedgerEntriesByEvent(ctx, eventType, eventID)
}

func (s *LedgerService) ListByService(ctx context.Context, sourceService string) ([]db.LedgerEntry, error) {
	return s.repo.ListLedgerEntriesByService(ctx, sourceService)
}
