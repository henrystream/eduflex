package ledger

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type LedgerClient struct {
	BaseURL string
}

type LedgerEntryRequest struct {
	EventType     string             `json:"event_type"`
	EventID       pgtype.UUID        `json:"event_id"`
	SourceService string             `json:"source_service"`
	DebitAccount  string             `json:"debit_account"`
	CreditAccount string             `json:"credit_account"`
	Amount        pgtype.Numeric     `json:"amount"`
	Currency      string             `json:"currency"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}

func NewLedgerClient(baseURL string) *LedgerClient {
	return &LedgerClient{BaseURL: baseURL}
}

func (c *LedgerClient) CreateEntry(req LedgerEntryRequest) error {
	body, _ := json.Marshal(req)
	_, err := http.Post(c.BaseURL+"/ledger", "application/json", bytes.NewBuffer(body))
	return err
}
