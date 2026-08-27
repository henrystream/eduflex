package client

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type FinancingClient struct {
	BaseURL string
}

func NewFinancingClient(url string) *FinancingClient {
	return &FinancingClient{BaseURL: url}
}

type Agreement struct {
	ID         pgtype.UUID    `json:"id"`
	StudentID  pgtype.UUID    `json:"student_id"`
	SchoolID   pgtype.UUID    `json:"school_id"`
	Principal  pgtype.Numeric `json:"principal"`
	TermMonths int            `json:"term_months"`
}

func (c *FinancingClient) GetAgreement(studentID pgtype.UUID) ([]Agreement, error) {
	resp, err := http.Get(c.BaseURL + "/agreements?student_id=" + studentID.String())
	//resp, err := http.Get(c.BaseURL + "/agreements/" + id.String())
	if err != nil {
		return []Agreement{}, err
	}
	defer resp.Body.Close()

	var a []Agreement
	err = json.NewDecoder(resp.Body).Decode(&a)
	return a, err
}
