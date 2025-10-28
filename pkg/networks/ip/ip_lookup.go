package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

func GetIp(domain string) (string, error) {
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

func GetIpInfo(ip string) (models.IpInfo, error) {
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
