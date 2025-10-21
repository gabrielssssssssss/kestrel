package models

type Website struct {
	Target     string      `json:"target"`
	Subdomains []Subdomain `json:"subdomains"`
	CMS        string      `json:"cms"`
}

type Domain struct {
	Ip          string      `json:"ip"`
	Domain      string      `json:"domain"`
	Certificate Certificate `json:"certificate"`
	Whois       Whois       `json:"whois"`
	Links       []string    `json:"links"`
}

type Subdomain struct {
	Ip          string      `json:"ip"`
	Subdomain   string      `json:"subdomain"`
	Certificate Certificate `json:"certificate"`
	Whois       Whois       `json:"whois"`
	Links       []string    `json:"links"`
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
