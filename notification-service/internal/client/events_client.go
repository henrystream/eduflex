package client

import (
	"encoding/json"
	"net/http"
)

type EventsClient struct {
	BaseURL string
}

func NewEventsClient(url string) *EventsClient {
	return &EventsClient{BaseURL: url}
}

type DomainEvent struct {
	ID            string          `json:"id"`
	EventType     string          `json:"event_type"`
	SourceService string          `json:"source_service"`
	AggregateID   string          `json:"aggregate_id"`
	Payload       json.RawMessage `json:"payload"`
	OccurredAt    string          `json:"occurred_at"`
}

func (c *EventsClient) ListUnprocessed() ([]DomainEvent, error) {
	resp, err := http.Get(c.BaseURL + "/events/unprocessed")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []DomainEvent
	err = json.NewDecoder(resp.Body).Decode(&events)
	return events, err
}

func (c *EventsClient) MarkProcessed(id string) error {
	req, _ := http.NewRequest("POST", c.BaseURL+"/events/"+id+"/processed", nil)
	_, err := http.DefaultClient.Do(req)
	return err
}
