package twitter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/config"
)

type Oauth struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type Twitter struct {
	Data Profile `json:"data"`
}

type Profile struct {
	Id              string               `json:"id"`
	Username        string               `json:"username"`
	Name            string               `json:"name"`
	PublicMetrics   ProfilePublicMetrics `json:"public_metrics"`
	CreatedAt       string               `json:"created_at"`
	ProfileImageUrl string               `json:"profile_image_url"`
	Verified        bool                 `json:"verified"`
	VerifiedType    string               `json:"verified_type"`
	PinnedTweetId   string               `json:"pinned_tweet_id"`
	Description     string               `json:"description"`
	Protected       bool                 `json:"protected"`
}

type ProfilePublicMetrics struct {
	FollowersCount string `json:"followers_count"`
	FollowingCount string `json:"following_count"`
	TweetCount     string `json:"tweet_count"`
	ListedCount    string `json:"listed_count"`
	LikeCount      string `json:"like_count"`
	MediaCount     string `json:"media_count"`
}

func NewBearerToken() (string, error) {
	var payload Oauth

	auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s",
			config.GetConfig("TWITTER_API_KEY"),
			config.GetConfig("TWITTER_API_KEY_SECRET"),
		),
	))

	request, err := http.NewRequest("POST", "https://api.x.com/oauth2/token",
		bytes.NewBufferString("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	request.Header.Add("Authorization", "Basic "+auth)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return "", err
	}

	return payload.AccessToken, nil
}

func GetProfile(username string) (Profile, error) {
	var payload Profile

	token, err := NewBearerToken()
	if err != nil {
		return payload, err
	}

	url := fmt.Sprintf("https://api.x.com/2/users/by/username/%s?user.fields=created_at,description,entities,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,verified_type,withheld", username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return payload, err
	}
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return payload, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
