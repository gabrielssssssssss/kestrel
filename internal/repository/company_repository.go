package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gabrielssssssssss/kestrel/internal/models"
)

type CompanyRepository struct{}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

func (r *CompanyRepository) FetchRechercheEntreprise(sirene string) (models.Company, error) {
	var payload models.Company

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
