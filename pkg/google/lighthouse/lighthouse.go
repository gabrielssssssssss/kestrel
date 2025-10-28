package lighthouse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/config"
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

func Monitor(url string) (*models.Lighthouse, error) {
	apiKey := config.GetConfig("LIGHTHOUSE_API_KEY")
	apiURL := fmt.Sprintf(
		"https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=%s&strategy=mobile&category=PERFORMANCE&category=ACCESSIBILITY&category=BEST_PRACTICES&category=SEO&key=%s",
		url,
		apiKey,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	categories := result["lighthouseResult"].(map[string]interface{})["categories"].(map[string]interface{})
	lh := &models.Lighthouse{}
	if perf, ok := categories["performance"].(map[string]interface{}); ok {
		if score, ok := perf["score"].(float64); ok {
			lh.Perfomance = score
		}
	}
	if seo, ok := categories["seo"].(map[string]interface{}); ok {
		if score, ok := seo["score"].(float64); ok {
			lh.SEO = score
		}
	}
	if a11y, ok := categories["accessibility"].(map[string]interface{}); ok {
		if score, ok := a11y["score"].(float64); ok {
			lh.Accessibility = score
		}
	}
	if bp, ok := categories["best-practices"].(map[string]interface{}); ok {
		if score, ok := bp["score"].(float64); ok {
			lh.BestPractices = score
		}
	}

	return lh, nil
}
