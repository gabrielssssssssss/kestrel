package search_engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Title    string
	Redirect string
	Snippet  string
}

func GetRawHtml(domain string, query string) ([]io.Reader, error) {
	var (
		page      int
		bodyArray []io.Reader
	)

	for {
		page++
		form := fmt.Sprintf(`lui=francais&language=francais&query=site:%s %s&cat=web&page=%d`,
			url.QueryEscape(domain), url.QueryEscape(query), page)
		search := strings.NewReader(form)

		request, err := http.NewRequest("POST", "https://www.startpage.com/sp/search", search)
		if err != nil {
			return nil, err
		}

		request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Add("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
		request.Header.Add("Referer", "https://www.startpage.com/")
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Add("Origin", "https://www.startpage.com")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("Sec-Fetch-Dest", "document")
		request.Header.Add("Sec-Fetch-Mode", "navigate")
		request.Header.Add("Sec-Fetch-Site", "same-origin")
		request.Header.Add("Sec-Fetch-User", "?1")
		request.Header.Add("Priority", "u=0, i")
		request.Header.Add("TE", "trailers")

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(body))
		if !strings.Contains(string(body), "w-gl") {
			break
		}

		bodyArray = append(bodyArray, bytes.NewReader(body))
		time.Sleep(1 * time.Second)
	}
	return bodyArray, nil
}

func ParseHtml(bodyArray []io.Reader) (string, error) {
	var payload []Result
	for _, body := range bodyArray {
		doc, err := goquery.NewDocumentFromReader(body)
		if err != nil {
			return "", err
		}

		doc.Find(".w-gl").Find(".result").Each(func(i int, s *goquery.Selection) {
			title := s.Find(".wgl-title").Text()
			redirect, _ := s.Find("a").Attr("href")
			snippet := s.Find("p").Text()
			payload = append(payload, Result{
				Title:    title,
				Redirect: redirect,
				Snippet:  snippet,
			})
		})
	}
	value, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(value), nil
}
