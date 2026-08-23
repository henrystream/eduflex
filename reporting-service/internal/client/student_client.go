package client

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type StudentClient struct {
	BaseURL string
}

func NewStudentClient(url string) *StudentClient {
	return &StudentClient{BaseURL: url}
}

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Disbursement struct {
	ID       string         `json:"id"`
	SchoolID string         `json:"school_id"`
	Amount   pgtype.Numeric `json:"amount"`
}

type Payment struct {
	ID       string         `json:"id"`
	SchoolID string         `json:"school_id"`
	Amount   pgtype.Numeric `json:"amount"`
}

func (c *StudentClient) GetStudent(studentID string) (Student, error) {
	resp, err := http.Get(c.BaseURL + "/students/" + studentID)
	if err != nil {
		return Student{}, err
	}
	defer resp.Body.Close()

	var student Student
	err = json.NewDecoder(resp.Body).Decode(&student)
	return student, err
}

func (c *StudentClient) ListDisbursementsBySchool(schoolID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/disbursements?school_id=" + schoolID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var disbursements []Disbursement
	err = json.NewDecoder(resp.Body).Decode(&disbursements)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(disbursements))
	for i, d := range disbursements {
		amounts[i] = d.Amount
	}

	return amounts, nil
}

func (c *StudentClient) ListPaymentsBySchool(schoolID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/payments?school_id=" + schoolID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payments []Payment
	err = json.NewDecoder(resp.Body).Decode(&payments)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(payments))
	for i, p := range payments {
		amounts[i] = p.Amount
	}

	return amounts, nil
}

func (c *StudentClient) ListPaymentsByStudent(studentID string) ([]pgtype.Numeric, error) {
	resp, err := http.Get(c.BaseURL + "/payments?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payments []Payment
	err = json.NewDecoder(resp.Body).Decode(&payments)
	if err != nil {
		return nil, err
	}

	amounts := make([]pgtype.Numeric, len(payments))
	for i, p := range payments {
		amounts[i] = p.Amount
	}

	return amounts, nil
}
