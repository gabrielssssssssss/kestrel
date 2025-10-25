package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gabrielssssssssss/kestrel/internal/config"
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

type GoogleMapsStruct struct{}

func NewMappyRepository() *GoogleMapsStruct {
	return &GoogleMapsStruct{}
}

func (r *GoogleMapsStruct) FetchPlaceId(query string) (string, error) {
	var payload models.GoogleMapsPlaceId

	params := url.Values{}
	params.Add("input", query)
	params.Add("inputtype", "textquery")
	params.Add("fields", "place_id")
	params.Add("key", config.GetConfig("GOOGLE_MAPS_API_KEY"))

	url := fmt.Sprint("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?", params.Encode())

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return "", err
	}

	return payload.Candidates[0].PlaceId, nil
}

func (r *GoogleMapsStruct) FetchPlaceDetails(placeId string) (models.GoogleMaps, error) {
	var payload models.GoogleMaps

	params := url.Values{}
	params.Add("place_id", placeId)
	params.Add("fields", "name,formatted_address,formatted_phone_number,website,opening_hours,price_level,rating,reviews,user_ratings_total,photos,business_status")
	params.Add("language", "fr")
	params.Add("key", config.GetConfig("GOOGLE_MAPS_API_KEY"))

	url := fmt.Sprint("https://maps.googleapis.com/maps/api/place/details/json?", params.Encode())

	response, err := http.Get(url)
	if err != nil {
		return payload, nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, nil
	}

	err = json.Unmarshal([]byte(body), &payload)

	return payload, nil
}
