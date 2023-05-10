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
	re, err := regexp.Compile(`<title>(.*?)</title>`)
	if err != nil {
		return ""
	}
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return ""
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

func IsStaticFile(urlStr string) bool {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		// URL解析错误
		return false
	}

	// 获取URL的路径部分
	path := parsedURL.Path

	// 获取路径中的文件名或路径
	parts := strings.Split(path, "/")
	fileName := parts[len(parts)-1]

	// 获取文件扩展名
	ext := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])

	// 判断扩展名是否为静态文件类型
	staticFileExtensions := []string{
		"png", "jpg", "jpeg", "gif", "css", "js", "ico", "svg", "ttf", "woff", "woff2", "eot", "otf", "map", "txt", "html", "htm", "json", "xml",
	}
	for _, staticExt := range staticFileExtensions {
		if ext == staticExt {
			return true
		}
	}

	return false
}
