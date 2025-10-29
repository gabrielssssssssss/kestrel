package emails

import (
	"encoding/json"
	"io"
	"net/http"
)

type Leakcheck struct {
	Success bool     `json:"success"`
	Found   int      `json:"found"`
	Fields  []string `json:"fields"`
	Sources []struct {
		Name string `json:"name"`
		Date string `json:"date"`
	}
}

func IsBreached(email string) Leakcheck {
	var payload Leakcheck

	response, err := http.Get("https://leakcheck.io/api/public?check=" + email)
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
