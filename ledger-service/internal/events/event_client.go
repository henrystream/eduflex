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
	return &EventClient{BaseURL: baseURL, SourceService: sourceService}
}

func (c *EventClient) Publish(eventType string, aggregateID pgtype.UUID, occurredAt pgtype.Timestamptz, payload interface{}) error {
	req := PublishEventRequest{
		EventType:     eventType,
		SourceService: c.SourceService,
		AggregateID:   aggregateID,
		Payload:       payload,
		OccurredAt:    occurredAt,
	}

	body, _ := json.Marshal(req)
	_, err := http.Post(c.BaseURL+"/events", "application/json", bytes.NewBuffer(body))
	return err
}
