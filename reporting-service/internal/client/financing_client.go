package client

import (
	"encoding/json"
	"fmt"
	"io"
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

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("financing service returned status %d: %s", resp.StatusCode, string(body))
	}

	var agreements []Agreement
	err = json.NewDecoder(resp.Body).Decode(&agreements)
	if err != nil {
		// Return empty array on decode error
		return []Agreement{}, nil
	}
	return agreements, nil
}

func (c *FinancingClient) ListInstallmentsByStudent(studentID string) ([]Installment, []pgtype.Numeric, error) {
	var (
		amounts      []pgtype.Numeric
		installments []Installment
	)
	agreements, _ := c.ListAgreementsByStudent(studentID)

	for _, agr := range agreements {
		resp, err := http.Get(c.BaseURL + "/installments?financing_id=" + agr.ID) //requires fid
		if err != nil {
			return nil, nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, nil, fmt.Errorf("financing service returned status %d: %s", resp.StatusCode, string(body))
		}

		err = json.NewDecoder(resp.Body).Decode(&installments)
		if err != nil {
			// Return empty array on decode error
			return nil, []pgtype.Numeric{}, nil
		}

		amounts = make([]pgtype.Numeric, len(installments))
		for i, inst := range installments {
			amounts[i] = inst.Amount
		}
	}

	return installments, amounts, nil
}
