package domain

import (
	"encoding/json"
	"io"
	"net/http"
)

type MicrosoftTenants struct {
	State                  int    `json:"State"`
	UserState              int    `json:"UserState"`
	Login                  string `json:"Login"`
	NameSpaceType          string `json:"NameSpaceType"`
	DomainName             string `json:"DomainName"`
	CloudInstanceName      string `json:"CloudInstanceName"`
	CloudInstanceIssuerUri string `json:"CloudInstanceIssuerUri"`
}

func FetchMicrosoftTenants(domain string) MicrosoftTenants {
	var payload MicrosoftTenants

	response, err := http.Get("https://login.microsoftonline.com/getuserrealm.srf?login=" + domain)
	if err != nil {
		return payload
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return payload
	}

	return payload
}
