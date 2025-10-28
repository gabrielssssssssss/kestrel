package wappalyzer

import (
	"io"
	"net/http"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func GetTechnologies(domain string) (map[string]wappalyzer.AppInfo, error) {
	response, err := http.DefaultClient.Get("https://" + domain)
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(response.Body)
	client, err := wappalyzer.New()
	if err != nil {
		return nil, err
	}

	fingerprints := client.FingerprintWithInfo(response.Header, body)
	return fingerprints, nil
}
