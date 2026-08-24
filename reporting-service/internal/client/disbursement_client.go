package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type DisbursementClient struct {
	BaseURL string
}

func NewDisbursementClient(url string) *DisbursementClient {
	return &DisbursementClient{BaseURL: url}
}

type DisbursementRecord struct {
	ID       string         `json:"id"`
	SchoolID string         `json:"school_id"`
	Amount   pgtype.Numeric `json:"amount"`
}

func (c *DisbursementClient) ListBySchool(schoolID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/disbursements?school_id=" + schoolID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("disbursement service returned status %d", resp.StatusCode)
	}

	var disbursements []DisbursementRecord
	err = json.NewDecoder(resp.Body).Decode(&disbursements)
	if err != nil {
		// Return empty array on decode error
		return []pgtype.Numeric{}, nil
	}

	amounts := make([]pgtype.Numeric, len(disbursements))
	for i, d := range disbursements {
		amounts[i] = d.Amount
	}

	return amounts, nil
}
