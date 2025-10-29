package medias

import (
	"encoding/json"
	"io"
	"net/http"
)

type Instagram struct {
	Username              string         `json:"username"`
	Fullname              string         `json:"full_name"`
	Pronouns              []string       `json:"pronouns"`
	Biography             string         `json:"biography"`
	BioLinks              BioLinks       `json:"bio_links"`
	ExternalUrl           string         `json:"external_url"`
	Followed              EdgeFollowedBy `json:"edge_followed_by"`
	Follow                EdgeFollow     `json:"edge_follow"`
	HasArEffects          bool           `json:"has_ar_effects"`
	HasClips              bool           `json:"has_clips"`
	HasGuides             bool           `json:"has_guides"`
	HashChannel           bool           `json:"has_channel"`
	HiglighReelCount      int16          `json:"highlight_reel_count"`
	IsBusinessAccount     bool           `json:"is_business_account"`
	IsProfessionalAccount bool           `json:"is_professional_account"`
	IsJoinedRecently      bool           `json:"is_joined_recently"`
	IsPrivate             bool           `json:"is_private"`
	IsVerified            bool           `json:"is_verified"`
	ProfilPicUrl          string         `json:"profile_pic_url"`
}

type BioLinks struct {
	Title    string `json:"title"`
	LynxUrl  string `json:"lynx_url"`
	Url      string `json:"url"`
	LinkType string `json:"link_type"`
}

type EdgeFollowedBy struct {
	Count int64 `json:"count"`
}

type EdgeFollow struct {
	Count int64 `json:"count"`
}

func GetInstagramProfile(username string) Instagram {
	var payload Instagram

	request, err := http.NewRequest("GET", "https://www.instagram.com/api/v1/users/web_profile_info?username="+username, nil)
	if err != nil {
		return payload
	}

	request.Header.Add("Authority", "www.instagram.com")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Referer", "https://www.instagram.com/"+username)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
	request.Header.Add("x-asbd-id", "198387")
	request.Header.Add("x-csrftoken", "VUm8uVUz0h2Y2CO1SwGgVAG3jQixNBmg")
	request.Header.Add("x-ig-app-id", "936619743392459")
	request.Header.Add("x-ig-www-claim", "0")
	request.Header.Add("x-requested-with", "XMLHttpRequest")

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
