package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Microsoft struct {
	IfExistsResult int `json:"IfExistsResult"`
}

func FetchMicrosoft(emailOrPhone string) bool {
	var payload Microsoft

	data := fmt.Sprintf(`{"username":"%s","isOtherIdpSupported":true,"checkPhones":true,"isRemoteNGCSupported":true,"isCookieBannerShown":false,"isFidoSupported":true,"country":"FR","forceotclogin":false,"isExternalFederationDisallowed":false,"isRemoteConnectSupported":false,"federationFlags":0,"isSignup":false,"isAccessPassSupported":true,"isQrCodePinSupported":true}`, emailOrPhone)
	request, err := http.NewRequest("POST", "https://login.microsoftonline.com/common/GetCredentialType?mkt=fr", strings.NewReader(data))
	if err != nil {
		return false
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Referer", "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=4765445b-32c6-49b0-83e6-1d93765276ca&redirect_uri=https%3A%2F%2Fwww.office.com%2Flandingv2&response_type=code%20id_token&scope=openid%20profile%20https%3A%2F%2Fwww.office.com%2Fv2%2FOfficeHome.All&response_mode=form_post&nonce=638972969930116057.MjZmOWRjYmItMDFjYS00MDA3LTljYTMtZGQ4YjFjNTkzNzU4YzYxZjg2MjktYWZjOC00NGM0LWFjNDAtZjAzNWE3YjRiYWQ0&ui_locales=fr&mkt=fr&client-request-id=ff012659-bed2-493f-a282-51652e1104ef&state=XLBc2wPNnbG6q8oX6LW7h3qVFmGuADOXB6YJGb-F7hBgwTNSwV6NQk7pSaRagztJOhV6mhTeZRRYvbtT9e27XVrItqSVN-AaWFORgzO0csdAPT8pxhJl9-UpFn2UM06bjQAwn4ZrHSoWLnwwmyPZHYJ8cFvAeQghtPn8e9RuD2zthtYf2i5BO4tSKMOFu5r_YJfQ8ASrFPJTSjLDVDvRtD28eJ1-ZCn9FIGr-AhiR5QBKlMY8eoju11aijkfZZ8xjyRG7U1QoRtxC6mYQ1C1o8-1N94t9-b_WyVuIUdb5Ko&x-client-SKU=ID_NET8_0&x-client-ver=8.5.0.0")
	request.Header.Add("hpgid", "1104")
	request.Header.Add("hpgact", "1800")
	request.Header.Add("Content-type", "application/json; charset=utf-8")
	request.Header.Add("Origin", "https://login.microsoftonline.com")
	request.Header.Add("Cookie", "fpc=AhrvSwEbGrNKhStnt1zIVx4; stsservicecookie=estsfd; x-ms-gateway-slice=estsfd")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return false
	}

	if payload.IfExistsResult == 5 {
		return true
	}

	return false
}
