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
	ID         string `json:"id"`
	FacilityID string `json:"facility_id"`
	Amount     string `json:"amount"`
}

func (c *LoanClient) ListDrawdownsByFacility(facilityID string) ([]Drawdown, error) {
	resp, err := http.Get(c.BaseURL + "/drawdowns?facility_id=" + facilityID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dd []Drawdown
	err = json.NewDecoder(resp.Body).Decode(&dd)
	return dd, err
}
