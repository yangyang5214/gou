package urlutil

import (
	"net"
	"net/url"
	"strings"
)

// GetIpsByDomain is 通过域名解析 ips
func GetIpsByDomain(domain string) []string {
	if strings.HasPrefix(domain, "http") {
		u, _ := url.Parse(domain)
		domain = u.Host
	}
	var ips []string
	ns, err := net.LookupIP(domain)
	if err != nil {
		return ips
	}

	for _, item := range ns {
		ips = append(ips, item.String())
	}
	return ips
}
