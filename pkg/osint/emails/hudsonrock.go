package emails

import (
	"encoding/json"
	"io"
	"net/http"
)

type Hudsonrock struct {
	Stealers []struct {
		TotalCorporateServices int      `json:"total_corporate_services"`
		TotalUserServices      int      `json:"total_user_services"`
		DateCompromised        string   `json:"date_compromised"`
		ComputerName           string   `json:"computer_name"`
		OperatingSystem        string   `json:"operating_system"`
		MalwarePath            string   `json:"malware_path"`
		Antiviruses            []string `json:"antiviruses"`
		Ip                     string   `json:"ip"`
		TopPasswords           []string `json:"top_passwords"`
		TopLogins              []string `json:"top_logins"`
	} `json:"stealers"`
}

func IsInfoSteal(email string) Hudsonrock {
	var payload Hudsonrock

	response, err := http.Get("https://cavalier.hudsonrock.com/api/json/v2/osint-tools/search-by-email?email=" + email)
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
