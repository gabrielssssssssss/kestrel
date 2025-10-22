package repository

import (
	"fmt"

	"github.com/gabrielssssssssss/kestrel/pkg/duckduckgo"
)

type LinkedinRepository struct{}

func NewLinkedinRepository() *LinkedinRepository {
	return &LinkedinRepository{}
}

func (r *LinkedinRepository) FetchLinkedinCompany(company string) ([]duckduckgo.SearchResult, error) {
	var payload []duckduckgo.SearchResult

	query := fmt.Sprintf(`q=site:fr.linkedin.com/in "%s"`, company)
	body, err := duckduckgo.GetRawHTML(query)
	fmt.Println(body)
	if err != nil {
		return payload, err
	}

	profiles, err := duckduckgo.ParseHTMLToJson(body)
	if err != nil {
		return payload, err
	}

	for _, profile := range profiles {
		payload = append(payload, duckduckgo.SearchResult{
			Title:   profile.Title,
			Link:    profile.Link,
			Snippet: profile.Snippet,
		})
	}
	return payload, nil
}
