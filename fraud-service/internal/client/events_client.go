package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type EventsClient struct {
	BaseURL       string
	SourceService string
}

func NewEventsClient(url, source string) *EventsClient {
	return &EventsClient{BaseURL: url, SourceService: source}
}

type PublishEventRequest struct {
	EventType     string             `json:"event_type"`
	SourceService string             `json:"source_service"`
	AggregateID   pgtype.UUID        `json:"aggregate_id"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
	Payload       interface{}        `json:"payload"`
}

func (c *EventsClient) Publish(eventType string, aggregateID pgtype.UUID, occurredAt pgtype.Timestamptz, payload interface{}) error {
	req := PublishEventRequest{
		EventType:     eventType,
		SourceService: c.SourceService,
		AggregateID:   aggregateID,
		OccurredAt:    occurredAt,
		Payload:       payload,
	}

	body, _ := json.Marshal(req)
	_, err := http.Post(c.BaseURL+"/events", "application/json", bytes.NewBuffer(body))
	return err
}
