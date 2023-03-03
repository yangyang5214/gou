package urlutil

import (
	"net/url"
	"regexp"
	"strings"
)

type Schema = string

const (
	HTTP  Schema = "http"
	HTTPS Schema = "https"
)

// GetHtmlTitle is parse title from html page
func GetHtmlTitle(html string) string {
	r, _ := regexp.Compile("<title>([^<]{1,200})</title>")
	title := r.FindString(html)
	if title == "" {
		return ""
	}
	title = strings.Replace(title, "<title>", "", 1)
	title = strings.Replace(title, "</title>", "", 1)
	return title
}

// AppendSchema is add scheme for url
func AppendSchema(u string, s Schema) string {
	if !strings.HasPrefix(u, HTTP) {
		u = s + "://" + u
	}
	return u
}

// GetUrlSuffix is return url Suffix. url ResourceType
func GetUrlSuffix(urlPath string) string {
	index := strings.LastIndex(urlPath, ".")
	if index == -1 {
		return ""
	}
	suffix := urlPath[index+1:]
	return suffix
}

// Parse is extend url.Parse, AppendSchema default use http
func Parse(urlStr string) (*url.URL, error) {
	rawUrl := AppendSchema(urlStr, HTTP)
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// RemoveQuery is remove query
func RemoveQuery(u string) string {
	r := strings.Split(u, "?")
	if len(r) == 1 {
		r = strings.Split(u, "%3F")
	}
	return r[0]
}

// GetUrlHost is return url host. if error return ""
func GetUrlHost(u string) string {
	urlParsed, err := Parse(u)
	if err != nil {
		return ""
	}
	return urlParsed.Host
}
