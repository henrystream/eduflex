package client

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type SchoolClient struct {
	BaseURL string
}

func NewSchoolClient(url string) *SchoolClient {
	return &SchoolClient{BaseURL: url}
}

type School struct {
	ID           pgtype.UUID      `json:"id"`
	Name         string           `json:"name"`
	Address      pgtype.Text      `json:"address"`
	ContactEmail pgtype.Text      `json:"contact_email"`
	ContactPhone pgtype.Text      `json:"contact_phone"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	Tier         string           `json:"tier"`
}

func (c *SchoolClient) GetSchool(id pgtype.UUID) (School, error) {
	resp, err := http.Get(c.BaseURL + "/schools/" + id.String())
	if err != nil {

		return School{}, err
	}
	defer resp.Body.Close()

	var s School
	err = json.NewDecoder(resp.Body).Decode(&s)
	return s, err
}
