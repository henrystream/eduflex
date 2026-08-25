package service

import (
	"context"
	"encoding/json"

	db "github.com/henrystream/eduflex/events-service/db/sqlc"
	"github.com/henrystream/eduflex/events-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(r *repository.EventRepository) *EventService {
	return &EventService{repo: r}
}

type PublishEventRequest struct {
	EventType     string             `json:"event_type"`
	SourceService string             `json:"source_service"`
	AggregateID   pgtype.UUID        `json:"aggregate_id"`
	Payload       interface{}        `json:"payload"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}

func (s *EventService) Publish(ctx context.Context, req PublishEventRequest) (db.DomainEvent, error) {
	payloadBytes, _ := json.Marshal(req.Payload)

	return s.repo.CreateEvent(ctx, repository.CreateEventParams{
		EventType:     req.EventType,
		SourceService: req.SourceService,
		AggregateID:   req.AggregateID,
		Payload:       payloadBytes,
		OccurredAt:    req.OccurredAt,
	})
}

func (s *EventService) ListUnprocessed(ctx context.Context) ([]db.DomainEvent, error) {
	return s.repo.ListUnprocessed(ctx)
}

func (s *EventService) MarkProcessed(ctx context.Context, id pgtype.UUID) error {
	return s.repo.MarkProcessed(ctx, id)
}
