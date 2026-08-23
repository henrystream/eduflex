package repository

import (
	"context"

	db "github.com/henrystream/eduflex/ledger-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type LedgerRepository struct {
	queries *db.Queries
}

func NewLedgerRepository(q *db.Queries) *LedgerRepository {
	return &LedgerRepository{queries: q}
}

type CreateLedgerEntryParams struct {
	EventType     string
	EventID       pgtype.UUID
	SourceService string
	DebitAccount  string
	CreditAccount string
	Amount        pgtype.Numeric
	Currency      string
	OccurredAt    pgtype.Timestamptz
}

func (r *LedgerRepository) CreateLedgerEntry(ctx context.Context, p CreateLedgerEntryParams) (db.LedgerEntry, error) {
	return r.queries.CreateLedgerEntry(ctx, db.CreateLedgerEntryParams{
		EventType:     p.EventType,
		EventID:       p.EventID,
		SourceService: p.SourceService,
		DebitAccount:  p.DebitAccount,
		CreditAccount: p.CreditAccount,
		Amount:        p.Amount,
		Currency:      p.Currency,
		OccurredAt:    p.OccurredAt,
	})
}

func (r *LedgerRepository) ListLedgerEntries(ctx context.Context) ([]db.LedgerEntry, error) {
	return r.queries.ListLedgerEntries(ctx)
}

func (r *LedgerRepository) ListLedgerEntriesByEvent(ctx context.Context, eventType string, eventID pgtype.UUID) ([]db.LedgerEntry, error) {
	return r.queries.ListLedgerEntriesByEvent(ctx, db.ListLedgerEntriesByEventParams{
		EventType: eventType,
		EventID:   eventID,
	})
}

func (r *LedgerRepository) ListLedgerEntriesByService(ctx context.Context, sourceService string) ([]db.LedgerEntry, error) {
	return r.queries.ListLedgerEntriesByService(ctx, sourceService)
}
