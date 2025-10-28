package whois

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func WhoisInfo(domain string) (whoisparser.WhoisInfo, error) {
	var payload whoisparser.WhoisInfo

	response, err := whois.Whois(domain)
	if err != nil {
		return payload, err
	}

	result, err := whoisparser.Parse(response)
	if err != nil {
		return payload, err
	}

	return result, nil
}
