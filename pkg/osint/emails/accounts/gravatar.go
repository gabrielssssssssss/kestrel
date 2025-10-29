package accounts

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gravatar struct {
	Entry [0]struct {
		Url         string `json:"profileUrl"`
		Username    string `json:"preferredUsername"`
		DisplayName string `json:"displayName"`
		Photo       string `json:"thumbnailUrl"`
	} `json:"entry"`
}

func FetchGravatar(email string) Gravatar {
	var payload Gravatar

	hash := md5.Sum([]byte(email))
	url := fmt.Sprintf("https://fr.gravatar.com/%s.json", hash)

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
