package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gabrielssssssssss/kestrel/internal/config"
)

type Skype map[string]SkypeData

type SkypeData struct {
	UserProfiles []SkypeProfile `json:"userProfiles"`
}

type SkypeProfile struct {
	ImageUri          string `json:"imageUri"`
	Cid               string `json:"cid"`
	UserPrincipalName string `json:"userPrincipalName"`
	GivenName         string `json:"givenName"`
	Surname           string `json:"surname"`
	Type              string `json:"type"`
}

func FetchSkype(email string) Skype {
	var payload Skype

	data := fmt.Sprintf(`{"emails":["%s"],"phones":[]}`, email)
	request, err := http.NewRequest("POST", "https://teams.live.com/api/mt/beta/users/searchUsers?ggEnabled=true", strings.NewReader(data))
	if err != nil {
		return payload
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	request.Header.Add("Referer", "https://teams.live.com/v2/worker/precompiled-web-worker-3ec9961eacbef443.js")
	request.Header.Add("x-ms-client-caller", "SEARCH_SUGGESTIONS")
	request.Header.Add("x-ms-client-type", "cdlworker")
	request.Header.Add("x-ms-client-version", "1415/25101615441")
	request.Header.Add("x-ms-session-id", "ad0511e3-6f84-40c8-97e9-20dfce2ba980")
	request.Header.Add("x-skypetoken", config.GetConfig("SKYPE_TOKEN"))
	request.Header.Add("Origin", "https://teams.live.com")
	request.Header.Add("Sec-GPC", "1")
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("authorization", "Bearer "+config.GetConfig("SKYPE_AUTHORIZATION"))
	request.Header.Add("x-ms-request-id", "")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Priority", "u=4")
	request.Header.Add("TE", "trailers")

	response, err := http.DefaultClient.Do(request)
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
