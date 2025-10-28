package certificate

import (
	"crypto/tls"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
)

type Certificate struct {
	Issuer             string   `json:"issuer"`
	CommonName         string   `json:"common_name"`
	SerialNumber       string   `json:"serial_number"`
	IssuerCountry      string   `json:"issuer_country"`
	SubjectAltNames    []string `json:"subject_alt_names"`
	SignatureAlgorithm string   `json:"signature_algorithm"`
	Created            string   `json:"created"`
	Expiry             string   `json:"expiry"`
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

	return payload, nil
}
