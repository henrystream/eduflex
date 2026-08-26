package client

import (
	"encoding/json"
	"net/http"
)

type DisbursementClient struct {
	BaseURL string
}

func NewDisbursementClient(url string) *DisbursementClient {
	return &DisbursementClient{BaseURL: url}
}

type Disbursement struct {
	ID            string `json:"id"`
	SchoolID      string `json:"school_id"`
	InvoiceID     string `json:"invoice_id"`
	Amount        string `json:"amount"`
	DisbursedAt   string `json:"disbursed_at"`
	PaymentMethod string `json:"payment_method"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
}

func (c *DisbursementClient) ListDisbursementsBySchool(schoolID string) ([]Disbursement, error) {
	resp, err := http.Get(c.BaseURL + "/disbursements?school_id=" + schoolID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var d []Disbursement
	err = json.NewDecoder(resp.Body).Decode(&d)
	return d, err
}
