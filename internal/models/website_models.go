package models

type Website struct {
	Target Domain `json:"target"`
}

type Domain struct {
	Ip           IpInfo            `json:"ip"`
	Domain       string            `json:"domain"`
	Ports        map[string]string `json:"ports"`
	Headers      map[string]string `json:"headers"`
	Certificate  Certificate       `json:"certificate"`
	Whois        Whois             `json:"whois"`
	Technologies []string          `json:"technologies"`
	Perfomances  Lighthouse        `json:"performances"`
	Links        []string          `json:"links"`
	Emails       []string          `json:"emails"`
	Phones       []string          `json:"phones"`
	Subdomains   []Subdomain       `json:"subdomains"`
}

type Subdomain struct {
	Ip           IpInfo              `json:"ip"`
	Subdomain    string              `json:"subdomain"`
	Headers      map[string]string   `json:"headers"`
	Certificate  Certificate         `json:"certificate"`
	Whois        Whois               `json:"whois"`
	Technologies map[string]struct{} `json:"technologies"`
	Links        []string            `json:"links"`
	Emails       []string            `json:"emails"`
	Phones       []string            `json:"phones"`
}

type IpInfo struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type Lighthouse struct {
	Perfomance    string `json:"perfomance"`
	Accessibility string `json:"accessibility"`
	BestPractices string `json:"best_practices"`
	SEO           string `json:"seo"`
}

type Whois struct {
	DomainName            string   `json:"domain_name"`
	Registrar             string   `json:"registrar"`
	RegistrarURL          string   `json:"registrar_url"`
	RegistrarIANAID       string   `json:"registrar_iana_id"`
	Creation              string   `json:"creation"`
	Updated               string   `json:"updated"`
	Expiry                string   `json:"expiry"`
	NameServers           []string `json:"name_servers"`
	AdminEmail            string   `json:"admin_email"`
	DomainStatus          string   `json:"domain_status"`
	CountryCode           string   `json:"country_code"`
	RegistrarAbuseContact string   `json:"registrar_abuse_contact"`
	DNSSEC                string   `json:"dnssec"`
}

type Certificate struct {
	CommonName        string   `json:"common_name"`
	SubjectAltNames   []string `json:"subject_alt_names"`
	Issuer            string   `json:"issuer"`
	SerialNumber      string   `json:"serial_number"`
	NotBefore         string   `json:"not_before"`
	NotAfter          string   `json:"not_after"`
	SignatureAlgo     string   `json:"signature_algo"`
	PublicKeyAlgo     string   `json:"public_key_algo"`
	FingerprintSHA256 string   `json:"fingerprint_sha256"`
	IssuerCountry     string   `json:"issuer_country"`
}
