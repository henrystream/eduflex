package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type EventsClient struct {
	BaseURL       string
	SourceService string
}

func NewEventsClient(url, source string) *EventsClient {
	return &EventsClient{BaseURL: url, SourceService: source}
}

type PublishEventRequest struct {
	EventType     string      `json:"event_type"`
	SourceService string      `json:"source_service"`
	AggregateID   string      `json:"aggregate_id"`
	Payload       interface{} `json:"payload"`
	OccurredAt    string      `json:"occurred_at"`
}

func (c *EventsClient) Publish(eventType, aggregateID, occurredAt string, payload interface{}) error {
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
