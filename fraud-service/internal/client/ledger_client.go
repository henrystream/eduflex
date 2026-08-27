package client

import (
	"encoding/json"
	"net/http"
)

type LedgerClient struct {
	BaseURL string
}

func NewLedgerClient(url string) *LedgerClient {
	return &LedgerClient{BaseURL: url}
}

type LedgerEntry struct {
	ID        string `json:"id"`
	EventType string `json:"event_type"`
	EventID   string `json:"event_id"`
	Amount    string `json:"amount"`
}

func (c *LedgerClient) ListByEvent(eventType, eventID string) ([]LedgerEntry, error) {
	resp, err := http.Get(c.BaseURL + "/ledger/by-event?event_type=" + eventType + "&event_id=" + eventID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []LedgerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	return entries, err
}
