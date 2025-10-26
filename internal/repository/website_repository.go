package repository

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/pkg/lighthouse"
	"github.com/gabrielssssssssss/kestrel/pkg/networks"
)

type WebsiteRepository struct{}

func NewWebsiteRepository() *WebsiteRepository {
	return &WebsiteRepository{}
}

func (r *WebsiteRepository) FetchWebsite(domain string) (models.Website, error) {
	var payload models.Website

	ipAddress, err := networks.GetIpFromDomain(domain)
	if err != nil {
		return payload, err
	}

	ipInfo, err := networks.GetIpInfoFromIp(ipAddress)
	if err != nil {
		return payload, err
	}

	// ports, err := networks.PortsOpen(ipAddress)
	// if err != nil {
	// 	return payload, err
	// }

	certificate, err := networks.GetSslCertificate(domain)
	if err != nil {
		return payload, err
	}

	// whois, err := networks.WhoisInfo(domain)
	// if err != nil {
	// 	return payload, err
	// }

	technologies, err := networks.GetTechnologies(domain)
	if err != nil {
		return payload, err
	}

	lighthouse, err := lighthouse.Monitor("https://" + domain)
	if err != nil {
		return payload, err
	}

	emails, phone, socialsMedias, err := networks.GetCoordinates(domain)
	if err != nil {
		return payload, err
	}

	subdomains, err := networks.GetSubdomain(domain)
	if err != nil {
		return payload, err
	}

	payload = models.Website{
		Target: models.Domain{
			Ip:     ipInfo,
			Domain: domain,
			// Ports:        ports,
			Certificate: certificate,
			// Whois:        whois,
			Technologies: technologies,
			Perfomances:  *lighthouse,
			Emails:       emails,
			Phones:       phone,
			SocialMedias: socialsMedias,
			Subdomains:   subdomains,
		},
	}

	return payload, nil
}
