package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AmazonToken struct {
	SessionToken string `json:"sessionToken"`
}

func RenewToken() string {
	var payload AmazonToken

	url := "https://www.amazon.com/aaut/verify/ap?options=%7B%22clientData%22%3A%22%7B%5C%22sessionId%5C%22%3A%5C%22146-7611126-8965132%5C%22%2C%5C%22marketplaceId%5C%22%3A%5C%22ATVPDKIKX0DER%5C%22%2C%5C%22rid%5C%22%3A%5C%22JDNR8EQKM3V5M6HXKWD2%5C%22%2C%5C%22ubid%5C%22%3A%5C%22135-5191116-8514348%5C%22%2C%5C%22pageType%5C%22%3A%5C%22AuthenticationPortal%5C%22%2C%5C%22appAction%5C%22%3A%5C%22SIGNIN_PWD_COLLECT%5C%22%2C%5C%22subPageType%5C%22%3A%5C%22SignInClaimCollect%5C%22%7D%22%2C%22challengeType%22%3Anull%2C%22locale%22%3A%22en-US%22%2C%22externalId%22%3Anull%2C%22enableHeaderFooter%22%3Atrue%2C%22enableBypassMechanism%22%3Afalse%2C%22enableModalView%22%3Afalse%2C%22eventTrigger%22%3A%22PageLoad%22%2C%22aaExternalToken%22%3Anull%2C%22forceJsFlush%22%3Atrue%2C%22aamationToken%22%3Anull%7D"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:121.0) Gecko/20100101 Firefox/121.0")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Referer", "Mozilla/5.0 (Windows NT 10.0; rv:121.0) Gecko/20100101 Firefox/121.0")
	request.Header.Add("User-Agent", "https://www.amazon.com/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.com%2F%3Fref_%3Dnav_ya_signin&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=usflex&openid.mode=checkid_setup&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ""
	}

	token := response.Header.Get("amz-aamation-resp")
	if token != "" {
		json.Unmarshal([]byte(token), &payload)
	}
	return payload.SessionToken
}

func FetchAmazon(emailOrPhone string) bool {
	token := RenewToken()

	form := url.Values{}
	form.Add("appActionToken", "oagjo32vt6fcLJoZjPejSjchaB0sVTR8mbQ4yMUjC7w=:2")
	form.Add("appAction", "SIGNIN_PWD_COLLECT")
	form.Add("subPageType", "SignInClaimCollect")
	form.Add("openid.return_to", "ape:aHR0cHM6Ly93d3cuYW1hem9uLmNvbS8/cmVmXz1uYXZfeWFfc2lnbmlu")
	form.Add("prevRID", "ape:SkROUjhFUUtNM1Y1TTZIWEtXRDI=")
	form.Add("workflowState", "eyJ6aXAiOiJERUYiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiQTI1NktXIn0.r7a-6gfxvMz9jy-dgtHj-LH94QAFxKo_CGcCGy5Ma7oMMhsclxiVXg.Lpo-WEliULY2j4Q8.Xmgkgg_1kvJcnsPwFKso2VJv-W5Xdomd6Ros3a6iqloJu2D6g2V5ilpDTRxsCB4zpA4yJvarLmflgoq6Kz4E_ma17XjR8U0-SlFxyhqNK07FJqNRMEMjeMQ8Xxsvr9FxhrWZ7eHJHYU0xa1ZOa0A5ahMEdynNE_5JfwCeU7JQ_sXN-4SNOXt5HNSlTG3TXrJOUik9Xg1OaViLYuVy8pnTQHdiFHzzFYLcB7GO6XiI30bcF2BZwKbGwG1-QW2TKHq_Q99402_lCYrwIt2onpMZhb6qpKHsyb9G0Ln4lw9nrFpBIzFcg.pNPH16elRES-C8z-ZSV48Q")
	form.Add("email", emailOrPhone)
	form.Add("password", "")
	form.Add("create", "0")
	form.Add("metadata1", "")
	form.Add("aaToken", fmt.Sprintf("{\"uniqueValidationId\":\"%s\"}", token))

	request, err := http.NewRequest("POST", "https://www.amazon.com/ap/signin", strings.NewReader(form.Encode()))
	if err != nil {
		return false
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:121.0) Gecko/20100101 Firefox/121.0")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", "https://www.amazon.com")
	request.Header.Add("DNT", "1")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Referer", "https://www.amazon.com/ap/signin?openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.com%2F%3Fref_%3Dnav_ya_signin&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=usflex&openid.mode=checkid_setup&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0")
	request.Header.Add("Cookie", "session-id=146-7611126-8965132; session-id-time=2379343441l; i18n-prefs=USD; lc-main=en_US; sp-cdn=\"L5Z9:FR\"; skin=noskin; csm-hit=tb:JDNR8EQKM3V5M6HXKWD2+s-NQ3WJXXCSB8H500XY65V|1748623498970&t:1748623498970&adb:adblk_no; ubid-main=135-5191116-8514348; session-token=JLdFssti4/YOtQlQkgkL7bD0Q7A6Y/D1Cr2SsRN+cBBMJPlBOJCKvN8YJJfRAu7a7KNfCg3I7daXR5dnYnVZbD3bUPmOEhyEz5hqBjygaCbbjb2dOFxvC9EF/SEepKJyzyEjg9dm5H6JupKN/XqaQX40bdWLnn6MLNRq5y2YjK03aLNQOW4wp/0yWTyICLqDqliQeSS4lYLqOvG83zG0kj4VBKF+bOxwND466thEeXcrIphZMtKVQzCwzH87PiuM539F+/xRHoQQrU0zkbbZBC66/wFyaqgSFutjb1YahoGjR+UIGn0IfrBNq79u9wa9KEURFlOGO+ymEzV/CVNKHOkrqodA7s+6; id_pkel=n0; id_pk=eyJuIjoiMCJ9")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("Sec-Fetch-Dest", "document")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("TE", "trailers")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	if strings.Contains((string(body)), "ap_password") {
		return true
	}
	return false
}
