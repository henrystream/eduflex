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
	ID        string         `json:"id"`
	StudentID string         `json:"student_id"`
	Principal pgtype.Numeric `json:"principal"`
	Term      int            `json:"term"`
}

type Installment struct {
	ID        string         `json:"id"`
	StudentID string         `json:"student_id"`
	Amount    pgtype.Numeric `json:"amount"`
	DueDate   string         `json:"due_date"`
}

func (c *FinancingClient) ListAgreementsByStudent(studentID string) ([]Agreement, error) {
	resp, err := http.Get(c.BaseURL + "/agreements?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var agreements []Agreement
	err = json.NewDecoder(resp.Body).Decode(&agreements)
	return agreements, err
}

func (c *FinancingClient) ListInstallmentsByStudent(studentID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/installments?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var installments []Installment
	err = json.NewDecoder(resp.Body).Decode(&installments)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(installments))
	for i, inst := range installments {
		amounts[i] = inst.Amount
	}

	return amounts, nil
}
