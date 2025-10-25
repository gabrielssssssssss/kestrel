package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gabrielssssssssss/kestrel/internal/config"
)

type Instagram struct {
	Data InstagramData `json:"data"`
}

type InstagramData struct {
	User InstagramUser `json:"user"`
}

type InstagramUser struct {
	IsMemorialized        bool     `json:"is_memorialized"`
	IsPrivate             bool     `json:"is_private"`
	Username              string   `json:"username"`
	ProfilePicUrl         string   `json:"profile_pic_url"`
	LatestReelMedia       int64    `json:"latest_reel_media"`
	Biography             string   `json:"biography"`
	Fullname              string   `json:"fullname"`
	IsVerified            bool     `json:"is_verified"`
	FollowerCount         int32    `json:"follower_count"`
	AddressStreet         string   `json:"address_street"`
	CityName              string   `json:"city_name"`
	IsBusiness            bool     `json:"is_business"`
	Zip                   string   `json:"zip"`
	ExternalUrl           string   `json:"external_url"`
	Pronouns              []string `json:"pronouns"`
	IsProfessionalAccount string   `json:"is_professional_account"`
	FollowingCount        int32    `json:"following_count"`
	MediaCount            int32    `json:"media_count"`
	TotalClipsCount       int32    `json:"total_clips_count"`
}

func GetProfileId(profilUrl string) (string, error) {
	request, err := http.NewRequest("GET", profilUrl, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Connection", "keep-alive")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	regex, err := regexp.Compile(`\bprofilePage_\d+\b`)
	if err != nil {
		return "", err
	}

	profilePage := regex.FindString(string(body))
	profileId := strings.Split(profilePage, "_")
	return profileId[1], nil
}

func GetProfile(profileId string) (Instagram, error) {
	var payload Instagram

	variables := fmt.Sprintf(`{"enable_integrity_filters":true,"id":"%s","render_surface":"PROFILE","__relay_internal__pv__PolarisProjectCannesEnabledrelayprovider":true,"__relay_internal__pv__PolarisProjectCannesLoggedInEnabledrelayprovider":true,"__relay_internal__pv__PolarisCannesGuardianExperienceEnabledrelayprovider":true,"__relay_internal__pv__PolarisCASB976ProfileEnabledrelayprovider":false,"__relay_internal__pv__PolarisRepostsConsumptionEnabledrelayprovider":false}`, profileId)

	data := fmt.Sprintf(
		"av=17841477772467953&__d=www&__user=0&__a=1&__req=1h&__hs=20386.HCSV2:instagram_web_pkg.2.1...0&dpr=3&__ccg=EXCELLENT&__rev=1028944355&__s=d52nrp:cr6kty:lp9vbh&__hsi=7565282531703529761&__dyn=7xeUjG1mxu1syUbFp41twpUnwgU7SbzEdF8aUco2qwJxS0k24o0B-q1ew6ywaq0yE462mcw5Mx62G5UswoEcE7O2l0Fwqo31w9O1lwxwQzXwae4UaEW2G0AEco5G0zK5o4q3y1Sw62wLyES1TwTU9UaQ0Lo6-3u2WE5B08-269wr86C1mgcEed6goK2O4UrAwCAxW1oxe6UaUaE4e1tyVrx60gm1hyEcE4ei16zomw&__csr=iM9kcgRrsW4hcy7MN6YaiRnvniLjibYjb9iQgPRWBKXWVKppptrzqqizpCCJaqleSGl7-h4FeAF6qAJ6GG_WGAbx7Cgig9VkWBBUGuqeAQ8AGGGV8C4aFxt6zbyoCGGUtxaQh3GjLxBHQibHGdppUB3p4U-aUGagByeGiy9UGiqVqy8e9Eiy-byVemi8y4Fo4Ocw05xWDlwee69p822w7-K1jS0k3Jprx9xrg0GS0LAfwxw5kwuE0Fi0ifw3Jk085xrg2GyVXxizU2mUeE9O5wam2de3106zm0i-0jyF2w5jCsE3VwaIUS3ng0ZW2G2pcbowk0qai01wRw1mS2S8U0q2wAw0zWwPw&__hsdp=geR0Fg8wmQyiNIclVzy5l8BQatwE84TdhIhKZ8-y90zwCwCbwNgaotzoN1xxryApwDyErgG689O0SgB0Oo422-dCgN169w_wwel0j82zwLwOK1fwGwNwj8K5o2iUlwxG3yGxnxCQ3y2m2K2q6o-1bw2484m0GU21wqE7i2l030UtwZwmu08xwZwlo1i82Vwc10eu7y0Gw8u3i&__hblp=0uo4S221wK5UO1bKqXzoC3m59o23wtE888FbzEG2a7E52tLCKUuGEO4FUd9pQ7UCEa8iga8b8oyU4e58cEnxi5WwKKUkxW6Uae8AxWi2O7ubxm0AC5o8qxS6GG5u7K3q2m2K2q6oRwzBwxw4awkE2OwxK7U2HG3u16wEwgo8o5a2l0cu2W1Yw5qwQxSu2ycwmu0FE1uEfo5m0i-68C0xoao467E5ut0dCdCxagwaEmwrodQ2q0Iopyo&__sjsp=geR0Fg8wlli9b5l35uoUxli9t2Doa21dPkr4rLI-y90di7oScgcuahC2uaxJ08m3Vw&__comet_req=7&fb_dtsg=NAfu_BiK-C9S3dBecBp5C522nTTzF7fJB_o5hh8i861qHldydcj4vdg:17853599968089360:1761427529&jazoest=26121&lsd=a5wfrdai25_R8HMdmgYbgX&__spin_r=1028944355&__spin_b=trunk&__spin_t=1761429601&__crn=comet.igweb.PolarisPostRouteNext&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=PolarisProfilePageContentQuery&server_timestamps=true&variables=%s&doc_id=24963806849976236",
		url.QueryEscape(variables),
	)
	request, err := http.NewRequest("POST", "https://www.instagram.com/graphql/query", strings.NewReader(data))
	if err != nil {
		return payload, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("X-FB-Friendly-Name", "PolarisProfilePageContentQuery")
	request.Header.Add("X-BLOKS-VERSION-ID", "d472af6df5cc606197723ed51adaa0886f926161310654a7c93600790814eba5")
	request.Header.Add("X-CSRFToken", "bGnU9HZlGjRHn8hTGGktF4JD6cMYANbp")
	request.Header.Add("X-IG-App-ID", "1217981644879628")
	request.Header.Add("X-Root-Field-Name", "fetch__XDTUserDict")
	request.Header.Add("X-FB-LSD", "a5wfrdai25_R8HMdmgYbgX")
	request.Header.Add("X-ASBD-ID", "359341")
	request.Header.Add("Origin", "https://www.instagram.com")
	request.Header.Add("Alt-Used", "www.instagram.com")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Referer", "https://www.instagram.com/p/DQGl8t7DD-Y/?chaining=true")
	request.Header.Add("Cookie", config.GetConfig("INSTAGRAM_COOKIE"))
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Priority", "u=0")
	request.Header.Add("TE", "trailers")

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
		return payload, err
	}

	return payload, nil
}
