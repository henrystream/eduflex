package client

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type LoanClient struct {
	BaseURL string
}

func NewLoanClient(url string) *LoanClient {
	return &LoanClient{BaseURL: url}
}

type Loan struct {
	ID           string         `json:"id"`
	StudentID    string         `json:"student_id"`
	Principal    pgtype.Numeric `json:"principal"`
	InterestRate float64        `json:"interest_rate"`
	Term         int            `json:"term"`
	Status       string         `json:"status"`
}

type LoanRepayment struct {
	ID        string         `json:"id"`
	LoanID    string         `json:"loan_id"`
	StudentID string         `json:"student_id"`
	Amount    pgtype.Numeric `json:"amount"`
	DueDate   string         `json:"due_date"`
	PaidDate  string         `json:"paid_date"`
}

func (c *LoanClient) GetLoan(loanID string) (Loan, error) {
	resp, err := http.Get(c.BaseURL + "/loans/" + loanID)
	if err != nil {
		return Loan{}, err
	}
	defer resp.Body.Close()

	var loan Loan
	err = json.NewDecoder(resp.Body).Decode(&loan)
	return loan, err
}

func (c *LoanClient) ListLoansByStudent(studentID string) ([]Loan, error) {
	resp, err := http.Get(c.BaseURL + "/loans?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var loans []Loan
	err = json.NewDecoder(resp.Body).Decode(&loans)
	return loans, err
}

func (c *LoanClient) ListRepaymentsByLoan(loanID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/loans/" + loanID + "/repayments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repayments []LoanRepayment
	err = json.NewDecoder(resp.Body).Decode(&repayments)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(repayments))
	for i, r := range repayments {
		amounts[i] = r.Amount
	}

	return amounts, nil
}

func (c *LoanClient) ListRepaymentsByStudent(studentID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/loans/repayments?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repayments []LoanRepayment
	err = json.NewDecoder(resp.Body).Decode(&repayments)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(repayments))
	for i, r := range repayments {
		amounts[i] = r.Amount
	}

	return amounts, nil
}
