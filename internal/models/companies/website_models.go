package models

import (
	whoisparser "github.com/likexian/whois-parser"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

type Website struct {
	Target Domain `json:"target"`
}

type Domain struct {
	Ip           IpInfo                        `json:"ip"`
	Domain       string                        `json:"domain"`
	Certificate  Certificate                   `json:"certificate"`
	Whois        whoisparser.WhoisInfo         `json:"whois"`
	Technologies map[string]wappalyzer.AppInfo `json:"technologies"`
	Perfomances  Lighthouse                    `json:"performances"`
	Emails       []string                      `json:"emails"`
	Phones       []string                      `json:"phones"`
	SocialMedias SocialMedias                  `json:"social_medias"`
	Subdomains   []string                      `json:"subdomains"`
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
	Perfomance    interface{} `json:"performance,omitempty"`
	Accessibility interface{} `json:"accessibility,omitempty"`
	BestPractices interface{} `json:"best_practices,omitempty"`
	SEO           interface{} `json:"seo,omitempty"`
}

type Certificate struct {
	Issuer             string `json:"issuer"`
	CommonName         string `json:"common_name"`
	SerialNumber       string `json:"serial_number"`
	IssuerCountry      string `json:"issuer_country"`
	SignatureAlgorithm string `json:"signature_algorithm"`
	Created            string `json:"created"`
	Expiry             string `json:"expiry"`
}

type SocialMedias struct {
	Facebook  []string `json:"facebook"`
	Instagram []string `json:"instagram"`
	Tiktok    []string `json:"tiktok"`
	Youtube   []string `json:"youtube"`
	Twitter   []string `json:"twitter"`
}
