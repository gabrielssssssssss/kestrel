package duckduckgo

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

func GetRawHTML(query string) (io.Reader, error) {
	payload := bytes.NewBuffer([]byte(query))

	request, err := http.NewRequest("POST", "https://html.duckduckgo.com/html/", payload)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Referer", "https://html.duckduckgo.com/")
	request.Header.Add("Origin", "https://html.duckduckgo.com")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Cookie", "kl=fr-fr")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("Sec-Fetch-Dest", "document")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-User", "?1")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Priority", "u=0, i")
	request.Header.Add("Pragma", "no-cache")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("TE", "trailers")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func ParseHTMLToJson(body io.Reader) ([]SearchResult, error) {
	var payload []SearchResult

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	doc.Find(".result").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h2 a").Text()
		link := s.Find(".result__extras__url a").Text()
		snippet := s.Find(".result__snippet").Text()
		if title != "" && link != "" && snippet != "" {
			payload = append(payload, SearchResult{
				Title:   strings.TrimSpace(title),
				Link:    strings.Join(strings.Fields(link), ""),
				Snippet: strings.TrimSpace(snippet),
			})
		}
	})

	return payload, nil
}
