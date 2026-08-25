package events

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type EventClient struct {
	BaseURL       string
	SourceService string
}

type PublishEventRequest struct {
	EventType     string             `json:"event_type"`
	SourceService string             `json:"source_service"`
	AggregateID   pgtype.UUID        `json:"aggregate_id"`
	Payload       interface{}        `json:"payload"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}

func NewEventClient(baseURL, sourceService string) *EventClient {
	return &EventClient{BaseURL: "http://events-service:8080", SourceService: sourceService}
}

func (c *EventClient) Publish(eventType, sourceService string, aggregateID pgtype.UUID, occurredAt pgtype.Timestamptz, payload interface{}) error {
	req := PublishEventRequest{
		EventType:     eventType,
		SourceService: sourceService,
		AggregateID:   aggregateID,
		Payload:       payload,
		OccurredAt:    occurredAt,
	}

	body, _ := json.Marshal(req)
	_, err := http.Post("http://events-service:8080/events", "application/json", bytes.NewBuffer(body))
	return err
}
