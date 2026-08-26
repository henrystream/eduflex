package client

import (
	"encoding/json"
	"net/http"
)

type StudentClient struct {
	BaseURL string
}

func NewStudentClient(url string) *StudentClient {
	return &StudentClient{BaseURL: url}
}

type StudentPayment struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	Amount    string `json:"amount"`
	PaidAt    string `json:"paid_at"`
	Reference string `json:"reference"`
}

func (c *StudentClient) ListPaymentsByStudent(studentID string) ([]StudentPayment, error) {
	resp, err := http.Get(c.BaseURL + "/payments?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payments []StudentPayment
	err = json.NewDecoder(resp.Body).Decode(&payments)
	return payments, err
}
