package linkpreview

import (
	"net/url"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseURL(urlStr string, baseURL string) string {
	if strings.HasPrefix(urlStr, "https://") || strings.HasPrefix(urlStr, "http://") {
		return urlStr
	}

	// URL is in relative form, use base url
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	resultURL := url.URL{
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   urlStr,
	}

	return resultURL.String()
}

func strOr(s1, s2 string) string {
	if s1 != "" {
		return s1
	}
	return s2
}
