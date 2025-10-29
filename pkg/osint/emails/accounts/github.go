package accounts

import (
	"encoding/json"
	"io"
	"net/http"
)

type Github struct {
	Items [0]struct {
		Id           int64  `json:"id"`
		Url          string `json:"htm_url"`
		Login        string `json:"login"`
		Photo        string `json:"avatar_url"`
		Type         string `json:"type"`
		UserViewType string `json:"user_view_type"`
	} `json:"items"`
}

func FetchGithub(email string) Github {
	var payload Github

	url := "https://api.github.com/search/users?q=" + email

	response, err := http.Get(url)
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
