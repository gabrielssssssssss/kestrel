package search_engine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResultItem struct {
	ClickTrackUrl       string `json:"clicktrackUrl"`
	ContentNoFormatting string `json:"contentNoFormatting"`
	TitleNoFormatting   string `json:"titleNoFormatting"`
	UnescapedUrl        string `json:"unescapedUrl"`
	RichSnippet         struct {
		Metatags struct {
			TwitterTitle     string `json:"twitterTitle"`
			TwitterImage     string `json:"twitterImage"`
			ProfileFirstName string `json:"profileFirstName"`
			ProfileLastName  string `json:"profileLastName"`
		} `json:"metatags"`
	} `json:"richSnippet"`
}

type GoogleCSEResponse struct {
	Results []ResultItem `json:"results"`
}

func GetRawHtml(domain string, query string) (string, error) {
	url := fmt.Sprintf(
		`https://cse.google.com/cse/element/v1?rsz=filtered_cse&num=100&hl=fr&source=gcsc&start=0&cselibv=6467658b9628de43&cx=f21a904778aed47f2&q=site:%s+%s&safe=off&cse_tok=AEXjvhJOpYaipm9pZuXADMljOpDn:1761604741534&lr=&cr=&gl=fr&filter=0&sort=&as_oq=&as_sitesearch=&exp=cc,apo&callback=google.search.cse.api18166&rurl=https://cse.google.com/cse?cx=f21a904778aed47f2&gl=fr&hl=fr&safe=off#gsc.tab=0&gsc.sort=&gsc.q=site%%3A%s%%20%s`,
		url.QueryEscape(domain),
		url.QueryEscape(query),
		url.QueryEscape(domain),
		url.QueryEscape(query),
	)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("Host", "cse.google.com")
	request.Header.Add("Sec-Ch-Ua-Full-Version-List", "")
	request.Header.Add("Sec-Ch-Ua-Platform", "\"Linux\"")
	request.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"141\", \"Not?A_Brand\";v=\"8\"")
	request.Header.Add("Sec-Ch-Ua-Model", "\"\"")
	request.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	request.Header.Add("Sec-Ch-Ua-Form-Factors", "")
	request.Header.Add("Sec-Ch-Ua-Wow64", "?0")
	request.Header.Add("Sec-Ch-Ua-Arch", "\"\"")
	request.Header.Add("Sec-Ch-Ua-Full-Version", "\"\"")
	request.Header.Add("Downlink", "10")
	request.Header.Add("Accept-Language", "fr-FR,fr;q=0.9")
	request.Header.Add("Sec-Ch-Prefers-Color-Scheme", "dark")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")
	request.Header.Add("Rtt", "150")
	request.Header.Add("Sec-Ch-Ua-Platform-Version", "\"\"")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("X-Client-Data", "CL6EywE=")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-Mode", "no-cors")
	request.Header.Add("Sec-Fetch-Dest", "script")
	request.Header.Add("Referer", "https://cse.google.com/cse?cx=f21a904778aed47f2&gl=fr&hl=fr&safe=off")
	request.Header.Add("Cookie", "AEC=AaJma5sFLV5T3oqQZVSmjQJoUTk9vQtESs7P8cUDxh7IMg108IbdnThHzxg; __Secure-ENID=29.SE=h1sSCxiGYeoTWW0ihzeCIn8LN3i6c0EatX_Gh0gE36lNoA4H8tIh19mkumcFe7IA23aMZCHzVJBv-Wskz8A5n_gatzjd29lK_HI7Jvr7UvYsOKuVtrUYf2SitJyD6sU8RillgarcFSqlETUHRY2dLgWzf_Ipeu9fN2QJxpv-Li66EJtgUIxo8O2lGygpMI7Snk1HZj43G7m9EUX4tsj5UA1gmQr0hoJelnRMn133O0DpG0OHTTjY3ms3Vw2S1B2aGAzF3xulN6fwriv1nMWUH6sG5NEf; __gsas=ID=b682f2f901bc7bed:T=1761604733:RT=1761604733:S=ALNI_MbvNal19-342m9rQWuWiXIUb5EtqA; __Secure-ENID=29.SE=dA6hcNFQKfJbDPjPswu_HWamNgi3DoI5V4IVTNOKfoDPAwKGMObO1TFeHVfFCr5Bl6JNFCkoEPtpGH9df4gffoXWSUnRpkoZeu71rA-1kjZYzrXfT_eg8nAx3mdyZaPwAEidZ7_rrzp1IH_8L7Z3XOyK1044mOgZHBGsUIowLlCi_t3XoIAH8ygB9GBR-RfJ1fFAqXQnGV_t_nO151mrdgkhkorfY5-PE-p92g98asRWLHc0RyJUWntogF2q")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ParseHtml(body string) (string, error) {
	var payload GoogleCSEResponse

	text := body[strings.Index(body, "(")+1 : strings.LastIndex(body, ")")]
	err := json.Unmarshal([]byte(text), &payload)
	if err != nil {
		return "", err
	}

	response, err := json.Marshal(payload.Results)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
