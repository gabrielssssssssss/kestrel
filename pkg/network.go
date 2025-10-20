package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

type IpInfo struct {
	IP       string `json:"ip"`
	HOSTNAME string `json:"hostname"`
	CITY     string `json:"city"`
	REGION   string `json:"region"`
	COUNTRY  string `json:"country"`
	LOC      string `json:"loc"`
	ORG      string `json:"org"`
	POSTAL   string `json:"postal"`
	TIMEZONE string `json:"timezone"`
}

func DomainToIp(domain string) string {
	ips, _ := net.LookupIP(domain)
	return ips[0].To4().String()
}

// func DomainToSubdomains(domain string) {
// 	url := fmt.Sprintf("https://sub-scan-api.reverseipdomain.com/?domain=%s", domain)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// }

func PortScanner(ip string) []string {
	var portsOpen []string
	for port := 0; port <= 65535; port++ {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, strconv.Itoa(port)), time.Duration(1)*time.Second)
		if err != nil {
			continue
		}
		defer conn.Close()
		portsOpen = append(portsOpen, strconv.Itoa(port))
	}
	return portsOpen
}

func IpGathering(ip string) (IpInfo, error) {
	var payload IpInfo

	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	resp, err := http.Get(url)
	if err != nil {
		return payload, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
