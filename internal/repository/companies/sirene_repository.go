package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/gabrielssssssssss/kestrel/internal/config"
)

type SireneRepository struct{}

func NewSireneRepository() *SireneRepository {
	return &SireneRepository{}
}

func (r *SireneRepository) FetchSirene(company, sector string) (string, error) {
	values := map[string]interface{}{
		"model": "gpt-5-mini-2025-08-07",
		"tools": []map[string]interface{}{
			{
				"type": "web_search",
				"user_location": map[string]interface{}{
					"type":    "approximate",
					"country": "FR",
				},
			},
		},
		"reasoning": map[string]interface{}{
			"effort": "low",
		},
		"input": fmt.Sprintf(`
			Tu es un professionnel de la recherche de numéro SIRENE français. 
			Tu dois répondre uniquement par le numéro SIRENE correspondant à l'entreprise mentionnée et rien d'autre. 
			Aucune explication, aucun texte supplémentaire, juste le numéro.
			Tu as le droit à deux informations, l'Entreprise et le Secteur.
			Entreprise: %s
			Secteur: %s
		`, company, sector),
	}

	body, err := json.Marshal(values)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	request.Header.Add("Authorization", "Bearer "+config.GetConfig("OPENAI_API_KEY"))
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	regex := regexp.MustCompile(`\b\d{9}\b`)
	return regex.FindString(string(body)), nil
}
