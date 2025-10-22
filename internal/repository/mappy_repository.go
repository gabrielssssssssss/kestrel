package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gabrielssssssssss/kestrel/internal/config"
	"github.com/gabrielssssssssss/kestrel/internal/models"
)

type MappyStruct struct{}

func NewMappyRepository() *MappyStruct {
	return &MappyStruct{}
}

func (r *MappyStruct) FetchMappySearch(query string) (models.MappySearch, error) {
	var payload models.MappySearch

	url := fmt.Sprintf("https://api-search.mappy.net/search/1.1/find?q=%s&f=places&bbox=46.74456,-1.70185,46.75386,-1.63962&extend_bbox=1&max_results=130&language=FRE&favorite_country=250&clientid=mappy&mid=559295169007&screensize=GE&tagid=SPD_RESPONSE_SEARCH", url.QueryEscape(query))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return payload, err
	}

	request.Header.Add("apikey", config.GetConfig("MAPPY_API_KEY"))
	request.Header.Add("Accept", "application/json, text/plain, */*")
	request.Header.Add("Referer", "https://fr.mappy.com/")
	request.Header.Add("Origin", "https://fr.mappy.com/")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return payload, err
	}
	defer response.Body.Close()

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

func (r *MappyStruct) FetchMappyGeo(id string) (models.MappyGeo, error) {
	var payload models.MappyGeo

	url := fmt.Sprintf("https://api-poi.mappy.net/data/poi/5.3/geoentity/%s.json?clientid=mappy&mid=517535225329&screensize=GE&abtest=roadbook_button-roadbook_etapes&tagid=SPD_DETAIL_GEOENTITY", id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return payload, nil
	}

	request.Header.Add("apikey", config.GetConfig("MAPPY_API_KEY"))
	request.Header.Add("Accept", "application/json, text/plain, */*")
	request.Header.Add("Referer", "https://fr.mappy.com/")
	request.Header.Add("Origin", "https://fr.mappy.com/")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return payload, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return payload, nil
	}

	return payload, nil
}
