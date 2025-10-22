package repository

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/pkg/networks"
)

type WebsiteRepository struct{}

func NewWebsiteRepository() *WebsiteRepository {
	return &WebsiteRepository{}
}

func (r *WebsiteRepository) FetchWebsite(domain string) (models.Website, error) {
	var payload models.Website

	// ipAddress, err := networks.GetIpFromDomain(domain)
	// if err != nil {
	// 	return payload, err
	// }
	// ipInfo, err := networks.GetIpInfoFromIp(ipAddress)
	// if err != nil {
	// 	return payload, err
	// }
	// ports := networks.PortsOpen(ipAddress)

	// url := fmt.Sprintf("http://%s", domain)
	// response, err := http.Get(url)
	// if err != nil {
	// 	return payload, err
	// }

	// headers := response.Header

	// fmt.Println(ipInfo)
	// fmt.Println(ports)
	// fmt.Println(headers)

	// truc, err := networks.GetSslCertificate(domain)
	// if err != nil {
	// 	return payload, err
	// }
	// fmt.Println(truc)
	networks.WhoisInfo(domain)
	return payload, nil
}
