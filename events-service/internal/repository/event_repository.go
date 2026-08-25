package repository

import (
	"context"

	db "github.com/henrystream/eduflex/events-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type EventRepository struct {
	queries *db.Queries
}

func NewEventRepository(q *db.Queries) *EventRepository {
	return &EventRepository{queries: q}
}

type CreateEventParams struct {
	EventType     string
	SourceService string
	AggregateID   pgtype.UUID
	Payload       []byte
	OccurredAt    pgtype.Timestamptz
}

func (r *EventRepository) CreateEvent(ctx context.Context, p CreateEventParams) (db.DomainEvent, error) {
	return r.queries.CreateEvent(ctx, db.CreateEventParams{
		EventType:     p.EventType,
		SourceService: p.SourceService,
		AggregateID:   p.AggregateID,
		Payload:       p.Payload,
		OccurredAt:    p.OccurredAt,
	})
}

func (r *EventRepository) ListUnprocessed(ctx context.Context) ([]db.DomainEvent, error) {
	return r.queries.ListUnprocessedEvents(ctx)
}

func (r *EventRepository) MarkProcessed(ctx context.Context, id pgtype.UUID) error {
	return r.queries.MarkEventProcessed(ctx, id)
}
