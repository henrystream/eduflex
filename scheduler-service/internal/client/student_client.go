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

type Payment struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

func (c *StudentClient) ListPaymentsByStudent(studentID string) ([]Payment, error) {
	resp, err := http.Get(c.BaseURL + "/payments?student_id=" + studentID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payments []Payment
	err = json.NewDecoder(resp.Body).Decode(&payments)
	return payments, err
}
