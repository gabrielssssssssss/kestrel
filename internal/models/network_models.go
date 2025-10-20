package models

type Network struct {
	DOMAIN      NetworkDomain `json:"domain"`
	PORTS       []string      `json:"ports"`
	DNS_RECORDS []string      `json:"dns_records"`
	CDN         []string      `json:"cdn"`
}

type NetworkDomain struct {
	IP         string              `json:"ip"`
	HOSTNAME   string              `json:"hostname"`
	PROVIDER   string              `json:"provider"`
	DOMAIN     string              `json:"domain"`
	ASN        string              `json:"asn"`
	LOCATION   string              `json:"location"`
	CREATED    string              `json:"created"`
	EXPIRE     string              `json:"expire"`
	UPDATED    string              `json:"updated"`
	SUBDOMAINS []NetworkSubdomains `json:"subdomains"`
}

type NetworkSubdomains struct {
	IP        string `json:"ip"`
	HOSTNAME  string `json:"hostname"`
	PROVIDER  string `json:"provider"`
	SUBDOMAIN string `json:"subdomain"`
	ASN       string `json:"asn"`
	LOCATION  string `json:"location"`
	TIMEZONE  string `json:"timezone"`
	CREATED   string `json:"created"`
	EXPIRE    string `json:"expire"`
	UPDATED   string `json:"updated"`
}
