package networks

import (
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/likexian/whois"
)

func GetIpFromDomain(domain string) (string, error) {
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

func GetIpInfoFromIp(ip string) (models.IpInfo, error) {
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

func GetPortService(port string) string {
	file, _ := os.Open("./data/ports_service.csv")
	lines := csv.NewReader(file)

	data, _ := lines.ReadAll()
	for _, row := range data {
		if row[0] == port {
			return row[1]
		}
	}
	return ""
}

func PortsOpen(host string) map[string]string {
	var (
		payload = make(map[string]string)
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)
	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, string(port)), 2*time.Second)
			if err != nil {
				return
			}
			if conn != nil {
				defer conn.Close()
				service := GetPortService(string(port))
				mutex.Lock()
				payload[port] = service
				mutex.Unlock()
			}
		}(strconv.Itoa(port))
	}
	wg.Wait()
	return payload
}

type Certificate struct {
	Issuer              string   `json:"issuer"`
	CommonName          string   `json:"common_name"`
	SerialNumber        string   `json:"serial_number"`
	IssuerCountry       string   `json:"issuer_country"`
	SubjectCountry      string   `json:"subject_country"`
	SubjectOrganization string   `json:"subject_organization"`
	SubjectAltNames     []string `json:"subject_alt_names"`
	SignatureAlgorithm  string   `json:"signature_algorithm"`
	Created             string   `json:"created"`
	Expiry              string   `json:"expiry"`
}

func GetSslCertificate(domain string) (models.Certificate, error) {
	var payload models.Certificate

	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return payload, err
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return payload, err
	}

	cert := state.PeerCertificates[0]

	payload = models.Certificate{
		Issuer:             cert.Issuer.CommonName,
		CommonName:         cert.Subject.CommonName,
		SerialNumber:       cert.SerialNumber.String(),
		SignatureAlgorithm: cert.SignatureAlgorithm.String(),
		Created:            cert.NotBefore.Format("2006-01-02"),
		Expiry:             cert.NotAfter.Format("2006-01-02"),
	}

	if len(cert.Issuer.Country) > 0 {
		payload.IssuerCountry = cert.Issuer.Country[0]
	}
	if len(cert.Subject.Country) > 0 {
		payload.SubjectCountry = cert.Subject.Country[0]
	}
	if len(cert.Subject.Organization) > 0 {
		payload.SubjectOrganization = cert.Subject.Organization[0]
	}

	return payload, nil
}

func WhoisInfo(domain string) (models.WhoisInfo, error) {
	var payload models.WhoisInfo
	result, err := whois.Whois(domain)
	if err == nil {
		fmt.Println(result)
	}
	return payload, nil
}
