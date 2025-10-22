package models

import "encoding/base64"

type Website struct {
	Target Domain `json:"target"`
}

type Domain struct {
	Ip           IpInfo            `json:"ip"`
	Domain       string            `json:"domain"`
	Ports        map[string]string `json:"ports"`
	Headers      map[string]string `json:"headers"`
	Certificate  Certificate       `json:"certificate"`
	Whois        WhoisInfo         `json:"whois"`
	Technologies []string          `json:"technologies"`
	Perfomances  Lighthouse        `json:"performances"`
	Links        []string          `json:"links"`
	Emails       []string          `json:"emails"`
	Phones       []string          `json:"phones"`
	Screenshot   base64.Encoding   `json:"screenshot"`
	Subdomains   []Subdomain       `json:"subdomains"`
}

type Subdomain struct {
	Ip           IpInfo              `json:"ip"`
	Subdomain    string              `json:"subdomain"`
	Headers      map[string]string   `json:"headers"`
	Certificate  Certificate         `json:"certificate"`
	Whois        WhoisInfo           `json:"whois"`
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

type WhoisInfo struct {
	Registrar                  string   `json:"registrar"`
	RegistrarURL               string   `json:"registrar_url"`
	RegistryDomainId           string   `json:"registry_domain_id"`
	RegistrarAbuseContactEmail string   `json:"registrar_abuse_contact_email"`
	RegistrarAbuseContactPhone string   `json:"registrar_abuse_contact_phone"`
	RegistrantName             string   `json:"registrant_name"`
	RegistrantOrganization     string   `json:"registrant_organization"`
	RegistrantStreet           string   `json:"registrant_street"`
	RegistrantCity             string   `json:"registrant_city"`
	RegistrantPostalCode       string   `json:"registrant_postal_code"`
	RegistrantCountry          string   `json:"registrant_country"`
	RegistrantPhone            string   `json:"registrant_phone"`
	RegistrantPhoneExt         string   `json:"registrant_phone_ext"`
	RegistrantFax              string   `json:"registrant_fax"`
	RegistrantFaxExt           string   `json:"registrant_fax_ext"`
	RegistrantEmail            string   `json:"registrant_email"`
	RegistrantAdminId          string   `json:"registrant_admin_id"`
	AdminName                  string   `json:"admin_name"`
	AdminOrganization          string   `json:"admin_organization"`
	AdminStreet                string   `json:"admin_street"`
	AdminCity                  string   `json:"admin_city"`
	AdminPostalCode            string   `json:"admin_postal_code"`
	AdminCountry               string   `json:"admin_country"`
	AdminPhone                 string   `json:"admin_phone"`
	AdminPhoneExt              string   `json:"admin_phone_ext"`
	AdminFax                   string   `json:"admin_fax"`
	AdminFaxExt                string   `json:"admin_fax_ext"`
	AdminEmail                 string   `json:"admin_email"`
	TechName                   string   `json:"techn_name"`
	TechOrganization           string   `json:"tech_organization"`
	TechStreet                 string   `json:"tech_street"`
	TechCity                   string   `json:"tech_city"`
	TechPostalCode             string   `json:"tech_postal_code"`
	TechCountry                string   `json:"tech_country"`
	TechPhone                  string   `json:"tech_phone"`
	TechPhoneExt               string   `json:"tech_phone_ext"`
	TechFax                    string   `json:"tech_fax"`
	TechFaxExt                 string   `json:"tech_fax_ext"`
	TechEmail                  string   `json:"tech_email"`
	BillingName                string   `json:"billing_name"`
	BillingOrganization        string   `json:"billing_organization"`
	BillingStreet              string   `json:"billing_street"`
	BillingCity                string   `json:"billing_city"`
	BillingPostalCode          string   `json:"billing_postal_code"`
	BillingCountry             string   `json:"billing_country"`
	BillingPhone               string   `json:"billing_phone"`
	BillingPhoneExt            string   `json:"billing_phone_ext"`
	BillingFax                 string   `json:"billing_fax"`
	BillingFaxExt              string   `json:"billing_fax_ext"`
	BillingEmail               string   `json:"billing_email"`
	NS                         []string `json:"names_servers"`
}

type Certificate struct {
	Issuer              string `json:"issuer"`
	CommonName          string `json:"common_name"`
	SerialNumber        string `json:"serial_number"`
	IssuerCountry       string `json:"issuer_country"`
	SubjectCountry      string `json:"subject_country"`
	SubjectOrganization string `json:"subject_organization"`
	SignatureAlgorithm  string `json:"signature_algorithm"`
	Created             string `json:"created"`
	Expiry              string `json:"expiry"`
}
