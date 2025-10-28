package coordinates

import (
	"io"
	"net/http"
	"regexp"
	"slices"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

const (
	emailExpression     = `(?:mailto:)?([a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,})`
	phoneExpression     = `tel:([+0-9][0-9 .\-()]{4,})`
	facebookExpression  = `https://(?:[a-zA-Z0-9-]+\.)?facebook\.com/[a-zA-Z0-9.]+/?`
	instagramExpression = `https://(?:[a-zA-Z0-9-]+\.)?(?:instagram\.com|instagr\.am)/[a-zA-Z0-9_.]+/?`
	tiktokExpression    = `https://(?:[a-zA-Z0-9-]+\.)?tiktok\.com/@[a-zA-Z0-9_.-]+/?`
	youtubeExpression   = `https://(?:[a-zA-Z0-9-]+\.)?youtube\.com/(?:channel|user)/[a-zA-Z0-9_-]+/?`
	twitterExpression   = `https://(?:[a-zA-Z0-9-]+\.)?twitter\.com/@?[a-zA-Z0-9_]+/?`
)

func GetCoordinates(domain string) ([]string, []string, models.SocialMedias, error) {
	var payload models.SocialMedias
	url := "https://" + domain

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, payload, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, payload, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, payload, err
	}
	content := string(body)

	var (
		regexEmail     = regexp.MustCompile(emailExpression)
		regexPhone     = regexp.MustCompile(phoneExpression)
		regexFacebook  = regexp.MustCompile(facebookExpression)
		regexInstagram = regexp.MustCompile(instagramExpression)
		regexTiktok    = regexp.MustCompile(tiktokExpression)
		regexYoutube   = regexp.MustCompile(youtubeExpression)
		regexTwitter   = regexp.MustCompile(twitterExpression)
	)

	emails := regexEmail.FindAllString(content, -1)
	phones := regexPhone.FindAllString(content, -1)
	fb := regexFacebook.FindAllString(content, -1)
	insta := regexInstagram.FindAllString(content, -1)
	tiktok := regexTiktok.FindAllString(content, -1)
	yt := regexYoutube.FindAllString(content, -1)
	twitter := regexTwitter.FindAllString(content, -1)

	emails = slices.Compact(emails)
	phones = slices.Compact(phones)
	payload = models.SocialMedias{
		Facebook:  slices.Compact(fb),
		Instagram: slices.Compact(insta),
		Tiktok:    slices.Compact(tiktok),
		Youtube:   slices.Compact(yt),
		Twitter:   slices.Compact(twitter),
	}

	return emails, phones, payload, nil
}
