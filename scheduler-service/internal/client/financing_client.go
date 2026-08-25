package client

import (
	"encoding/json"
	"net/http"
)

type FinancingClient struct {
	BaseURL string
}

func NewFinancingClient(url string) *FinancingClient {
	return &FinancingClient{BaseURL: url}
}

type Installment struct {
	ID        string `json:"id"`
	Amount    string `json:"amount"`
	DueDate   string `json:"due_date"`
	Status    string `json:"status"`
	Financing string `json:"financing_id"`
}

func (c *FinancingClient) ListInstallments(financingID string) ([]Installment, error) {
	resp, err := http.Get(c.BaseURL + "/installments?financing_id=" + financingID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var installments []Installment
	err = json.NewDecoder(resp.Body).Decode(&installments)
	return installments, err
}
