package repository

import (
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	"github.com/gabrielssssssssss/kestrel/pkg/google/lighthouse"
	cert "github.com/gabrielssssssssss/kestrel/pkg/networks/certificate"
	"github.com/gabrielssssssssss/kestrel/pkg/networks/coordinates"
	"github.com/gabrielssssssssss/kestrel/pkg/networks/ip"
	"github.com/gabrielssssssssss/kestrel/pkg/networks/subdomain"
	wp "github.com/gabrielssssssssss/kestrel/pkg/networks/wappalyzer"
	wh "github.com/gabrielssssssssss/kestrel/pkg/networks/whois"
)

type WebsiteRepository struct{}

func NewWebsiteRepository() *WebsiteRepository {
	return &WebsiteRepository{}
}

func (r *WebsiteRepository) FetchWebsite(domain string) (models.Website, error) {
	var payload models.Website

	ipAddress, err := ip.GetIp(domain)
	if err != nil {
		return payload, err
	}

	ipInfo, err := ip.GetIpInfo(ipAddress)
	if err != nil {
		return payload, err
	}

	certificate, err := cert.GetSslCertificate(domain)
	if err != nil {
		return payload, err
	}

	whois, err := wh.WhoisInfo(domain)
	if err != nil {
		return payload, err
	}

	technologies, err := wp.GetTechnologies(domain)
	if err != nil {
		return payload, err
	}

	lh, err := lighthouse.Monitor("https://" + domain)
	if err != nil {
		return payload, err
	}

	emails, phone, socialsMedias, err := coordinates.GetCoordinates(domain)
	if err != nil {
		return payload, err
	}

	subdomains, err := subdomain.GetSubdomain(domain)
	if err != nil {
		return payload, err
	}

	payload = models.Website{
		Target: models.Domain{
			Ip:           ipInfo,
			Domain:       domain,
			Certificate:  certificate,
			Whois:        whois,
			Technologies: technologies,
			Perfomances:  *lh,
			Emails:       emails,
			Phones:       phone,
			SocialMedias: socialsMedias,
			Subdomains:   subdomains,
		},
	}

	return payload, nil
}
