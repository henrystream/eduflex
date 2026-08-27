package fraud

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type FraudClient struct {
	BaseURL string
}

func NewFraudClient(url string) *FraudClient {
	return &FraudClient{BaseURL: url}
}

type FraudCheckResult struct {
	StudentRisk struct {
		Score float64 `json:"score"`
		Level string  `json:"level"`
	} `json:"student_risk"`

	SchoolRisk struct {
		Score float64 `json:"score"`
		Level string  `json:"level"`
	} `json:"school_risk"`

	AgreementRisk struct {
		Score float64 `json:"score"`
		Level string  `json:"level"`
	} `json:"agreement_risk"`
}

func (c *FraudClient) CheckAgreement(agreementID pgtype.UUID) (FraudCheckResult, error) {
	resp, err := http.Get(c.BaseURL + "/fraud/check-agreement?agreement_id=" + agreementID.String())
	if err != nil {
		return FraudCheckResult{}, err
	}
	defer resp.Body.Close()

	var result FraudCheckResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
