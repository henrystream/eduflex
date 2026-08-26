package client

import (
	"encoding/json"
	"net/http"
)

type LoanClient struct {
	BaseURL string
}

func NewLoanClient(url string) *LoanClient {
	return &LoanClient{BaseURL: url}
}

type Drawdown struct {
	ID           string `json:"id"`
	FacilityID   string `json:"facility_id"`
	Amount       string `json:"amount"`
	DrawdownDate string `json:"drawdown_date"`
	Reference    string `json:"reference"`
}

type Repayment struct {
	ID         string `json:"id"`
	DrawdownID string `json:"drawdown_id"`
	Amount     string `json:"amount"`
	PaidAt     string `json:"paid_at"`
	Reference  string `json:"reference"`
}

func (c *LoanClient) ListDrawdowns(facilityID string) ([]Drawdown, error) {
	resp, err := http.Get(c.BaseURL + "/drawdowns?facility_id=" + facilityID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dd []Drawdown
	err = json.NewDecoder(resp.Body).Decode(&dd)
	return dd, err
}

func (c *LoanClient) ListRepayments(drawdownID string) ([]Repayment, error) {
	resp, err := http.Get(c.BaseURL + "/repayments?drawdown_id=" + drawdownID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rp []Repayment
	err = json.NewDecoder(resp.Body).Decode(&rp)
	return rp, err
}
