package subdomain

import (
	"encoding/json"
	"io"
	"net/http"
)

type Subdomain struct {
	Result struct {
		Domains []string `json:"domains"`
	} `json:"result"`
}

func GetSubdomain(domain string) ([]string, error) {
	var payload Subdomain
	url := "https://sub-scan-api.reverseipdomain.com/?domain=" + domain

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return nil, err
	}

	return payload.Result.Domains, nil
}
