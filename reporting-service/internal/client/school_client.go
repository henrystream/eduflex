package client

import (
	"encoding/json"
	"net/http"
)

type SchoolClient struct {
	BaseURL string
}

func NewSchoolClient(url string) *SchoolClient {
	return &SchoolClient{BaseURL: url}
}

type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *SchoolClient) GetSchool(schoolID string) (School, error) {
	resp, err := http.Get(c.BaseURL + "/schools/" + schoolID)
	if err != nil {
		return School{}, err
	}
	defer resp.Body.Close()

	var school School
	err = json.NewDecoder(resp.Body).Decode(&school)
	return school, err
}
