package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

type OrganizationRepository struct{}

func NewCompanyRepository() *OrganizationRepository {
	return &OrganizationRepository{}
}

func (r *OrganizationRepository) FetchOrganization(sirene string) (models.Organization, error) {
	var payload models.Organization

	params := url.Values{}
	params.Add("q", sirene)

	response, err := http.Get("https://recherche-entreprises.api.gouv.fr/search?" + params.Encode())
	if err != nil {
		return payload, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
