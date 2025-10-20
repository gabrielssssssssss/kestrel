package repository

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/models"
)

func FetchRechercheEntreprise(url string) (models.Company, error) {
	var payload models.Company

	response, err := http.Get(url)
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
