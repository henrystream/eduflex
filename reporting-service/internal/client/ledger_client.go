package client

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type LedgerClient struct {
	BaseURL string
}

func NewLedgerClient(url string) *LedgerClient {
	return &LedgerClient{BaseURL: url}
}

func (c *LedgerClient) ListLedgerEntries() ([]LedgerEntry, error) {
	resp, err := http.Get(c.BaseURL + "/ledger")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []LedgerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	return entries, err
}

func (c *LedgerClient) ListByService(service string) ([]LedgerEntry, error) {
	resp, err := http.Get(c.BaseURL + "/ledger/by-service?source_service=" + service)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []LedgerEntry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	return entries, err
}

type LedgerEntry struct {
	ID            string             `json:"id"`
	EventType     string             `json:"event_type"`
	EventID       pgtype.UUID        `json:"event_id"`
	SourceService string             `json:"source_service"`
	DebitAccount  string             `json:"debit_account"`
	CreditAccount string             `json:"credit_account"`
	Amount        pgtype.Numeric     `json:"amount"`
	Currency      string             `json:"currency"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}
