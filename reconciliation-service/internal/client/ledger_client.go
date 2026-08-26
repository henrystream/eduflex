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
	ID            string `json:"id"`
	EventType     string `json:"event_type"`
	EventID       string `json:"event_id"`
	DebitAccount  string `json:"debit_account"`
	CreditAccount string `json:"credit_account"`
	Amount        string `json:"amount"`
	OccurredAt    string `json:"occurred_at"`
}

func (c *LedgerClient) ListAll() ([]LedgerEntry, error) {
	resp, err := http.Get(c.BaseURL + "/ledger")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []LedgerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	return entries, err
}
