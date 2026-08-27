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

type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	NationalID string `json:"national_id"`
	Country    string `json:"country"`
}

func (c *StudentClient) GetStudent(id string) (Student, error) {
	resp, err := http.Get(c.BaseURL + "/students/" + id)
	if err != nil {
		return Student{}, err
	}
	defer resp.Body.Close()

	var s Student
	err = json.NewDecoder(resp.Body).Decode(&s)
	return s, err
}
