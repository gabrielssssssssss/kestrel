package networks

import (
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/likexian/whois"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

type Subdomain struct {
	Result SubdomainResult `json:"result"`
}

type SubdomainResult struct {
	Domains []string `json:"domains"`
}

const (
	emailExpression     = `mailto:([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`
	phoneExpression     = `tel:([+0-9][0-9 .\-()]{4,})`
	facebookExpression  = `https://(?:[a-zA-Z0-9-]+\.)?facebook\.com/[a-zA-Z0-9.]+/?`
	instagramExpression = `https://(?:[a-zA-Z0-9-]+\.)?(?:instagram\.com|instagr\.am)/[a-zA-Z0-9_.]+/?`
	tiktokExpression    = `https://(?:[a-zA-Z0-9-]+\.)?tiktok\.com/@[a-zA-Z0-9_.-]+/?`
	youtubeExpression   = `https://(?:[a-zA-Z0-9-]+\.)?youtube\.com/(?:channel|user)/[a-zA-Z0-9_-]+/?`
	twitterExpression   = `https://(?:[a-zA-Z0-9-]+\.)?twitter\.com/@?[a-zA-Z0-9_]+/?`
)

func GetIpFromDomain(domain string) (string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", nil
}

func GetIpInfoFromIp(ip string) (models.IpInfo, error) {
	var payload models.IpInfo

	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	response, err := http.Get(url)
	if err != nil {
		return payload, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	return payload, err
}

func GetPortService(port string) string {
	file, _ := os.Open("./data/ports_service.csv")
	lines := csv.NewReader(file)

	data, _ := lines.ReadAll()
	for _, row := range data {
		if row[0] == port {
			return row[1]
		}
	}
	return ""
}

func PortsOpen(host string) (map[string]string, error) {
	var (
		payload = make(map[string]string)
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)
	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, string(port)), 2*time.Second)
			if err != nil {
				return
			}
			if conn != nil {
				defer conn.Close()
				service := GetPortService(string(port))
				mutex.Lock()
				payload[port] = service
				mutex.Unlock()
			}
		}(strconv.Itoa(port))
	}
	wg.Wait()
	return payload, nil
}

type Certificate struct {
	Issuer              string   `json:"issuer"`
	CommonName          string   `json:"common_name"`
	SerialNumber        string   `json:"serial_number"`
	IssuerCountry       string   `json:"issuer_country"`
	SubjectCountry      string   `json:"subject_country"`
	SubjectOrganization string   `json:"subject_organization"`
	SubjectAltNames     []string `json:"subject_alt_names"`
	SignatureAlgorithm  string   `json:"signature_algorithm"`
	Created             string   `json:"created"`
	Expiry              string   `json:"expiry"`
}

func GetSslCertificate(domain string) (models.Certificate, error) {
	var payload models.Certificate

	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return payload, err
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return payload, err
	}

	cert := state.PeerCertificates[0]

	payload = models.Certificate{
		Issuer:             cert.Issuer.CommonName,
		CommonName:         cert.Subject.CommonName,
		SerialNumber:       cert.SerialNumber.String(),
		SignatureAlgorithm: cert.SignatureAlgorithm.String(),
		Created:            cert.NotBefore.Format("2006-01-02"),
		Expiry:             cert.NotAfter.Format("2006-01-02"),
	}

	if len(cert.Issuer.Country) > 0 {
		payload.IssuerCountry = cert.Issuer.Country[0]
	}
	if len(cert.Subject.Country) > 0 {
		payload.SubjectCountry = cert.Subject.Country[0]
	}
	if len(cert.Subject.Organization) > 0 {
		payload.SubjectOrganization = cert.Subject.Organization[0]
	}

	return payload, nil
}

func antiDuplicate(value string) string {
	value = strings.TrimSpace(strings.ReplaceAll(value, "\t", " "))
	value = strings.Trim(value, ":")

	lower := strings.ToLower(value)
	if strings.Contains(lower, "redacted") || strings.Contains(lower, "privacy") || strings.Contains(lower, "protected") {
		return ""
	}

	parts := strings.Split(value, ",")
	unique := []string{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		found := false
		for _, u := range unique {
			if u == part {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, part)
		}
	}
	return strings.Join(unique, ", ")
}

func WhoisInfo(domain string) (models.WhoisInfo, error) {
	var payload models.WhoisInfo

	result, err := whois.Whois(domain)
	if err != nil {
		return payload, err
	}

	search := func(key string) string {
		re := regexp.MustCompile(fmt.Sprintf(`(?mi)^\s*%s\s*:\s*(.+)$`, regexp.QuoteMeta(key)))
		matches := re.FindAllStringSubmatch(result, -1)
		if len(matches) == 0 {
			return ""
		}

		values := []string{}
		for _, m := range matches {
			if len(m) > 1 {
				cleaned := antiDuplicate(m[1])
				if cleaned != "" {
					values = append(values, cleaned)
				}
			}
		}
		return antiDuplicate(strings.Join(values, ", "))
	}

	payload = models.WhoisInfo{
		Domain:                     search("Domain Name"),
		Registrar:                  search("Registrar"),
		Created:                    search("Creation Date"),
		ExpiryDate:                 search("Registry Expiry Date"),
		LastUpdate:                 search("Updated Date"),
		RegistrarURL:               search("Registrar URL"),
		RegistryDomainId:           search("Registry Domain ID"),
		RegistrarAbuseContactEmail: search("Registrar Abuse Contact Email"),
		RegistrarAbuseContactPhone: search("Registrar Abuse Contact Phone"),

		RegistrantName:         search("Registrant Name"),
		RegistrantOrganization: search("Registrant Organization"),
		RegistrantStreet:       search("Registrant Street"),
		RegistrantCity:         search("Registrant City"),
		RegistrantPostalCode:   search("Registrant Postal Code"),
		RegistrantCountry:      search("Registrant Country"),
		RegistrantPhone:        search("Registrant Phone"),
		RegistrantPhoneExt:     search("Registrant Phone Ext"),
		RegistrantFax:          search("Registrant Fax"),
		RegistrantFaxExt:       search("Registrant Fax Ext"),
		RegistrantEmail:        search("Registrant Email"),
		RegistrantAdminId:      search("Registry Registrant ID"),

		AdminName:         search("Admin Name"),
		AdminOrganization: search("Admin Organization"),
		AdminStreet:       search("Admin Street"),
		AdminCity:         search("Admin City"),
		AdminPostalCode:   search("Admin Postal Code"),
		AdminCountry:      search("Admin Country"),
		AdminPhone:        search("Admin Phone"),
		AdminPhoneExt:     search("Admin Phone Ext"),
		AdminFax:          search("Admin Fax"),
		AdminFaxExt:       search("Admin Fax Ext"),
		AdminEmail:        search("Admin Email"),

		TechName:         search("Tech Name"),
		TechOrganization: search("Tech Organization"),
		TechStreet:       search("Tech Street"),
		TechCity:         search("Tech City"),
		TechPostalCode:   search("Tech Postal Code"),
		TechCountry:      search("Tech Country"),
		TechPhone:        search("Tech Phone"),
		TechPhoneExt:     search("Tech Phone Ext"),
		TechFax:          search("Tech Fax"),
		TechFaxExt:       search("Tech Fax Ext"),
		TechEmail:        search("Tech Email"),

		BillingName:         search("Billing Name"),
		BillingOrganization: search("Billing Organization"),
		BillingStreet:       search("Billing Street"),
		BillingCity:         search("Billing City"),
		BillingPostalCode:   search("Billing Postal Code"),
		BillingCountry:      search("Billing Country"),
		BillingPhone:        search("Billing Phone"),
		BillingPhoneExt:     search("Billing Phone Ext"),
		BillingFax:          search("Billing Fax"),
		BillingFaxExt:       search("Billing Fax Ext"),
		BillingEmail:        search("Billing Email"),

		NS: search("Name Server"),
	}

	return payload, nil
}

func GetSubdomain(domain string) ([]string, error) {
	var payload Subdomain
	url := "https://sub-scan-api.reverseipdomain.com/?domain=" + domain

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return nil, err
	}

	return payload.Result.Domains, nil
}

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
