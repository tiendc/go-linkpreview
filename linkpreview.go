package linkpreview

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func validateURL(link string) error {
	parsedURL, err := url.ParseRequestURI(link)
	if err != nil || !strings.HasPrefix(parsedURL.Scheme, "http") {
		return errors.New("invalid URL")
	}
	return nil
}

func Parse(link string, options ...ConfigOption) (ParserResult, error) {
	if err := validateURL(link); err != nil {
		return ParserResult{}, err
	}

	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return ParserResult{}, err
	}

	// Add security headers
	request.Header.Set("X-Content-Type-Options", "nosniff")
	request.Header.Set("X-Frame-Options", "DENY")
	request.Header.Set("X-XSS-Protection", "1; mode=block")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ParserResult{}, err
	}
	defer response.Body.Close()

	return ParseFromReader(link, response.Body, options...)
}

func ParseFromReader(link string, data io.Reader, options ...ConfigOption) (ParserResult, error) {
	if err := validateURL(link); err != nil {
		return ParserResult{}, err
	}

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return ParserResult{}, err
	}

	ctx := ParserContext{
		Link: link,
		Config: ParserConfig{
			ParseOGMeta:      true,
			ParseTwitterMeta: true,
		},
		Doc: doc,
	}
	for _, opt := range options {
		opt(&ctx.Config)
	}

	err = ctx.Parse()
	if err != nil {
		return ParserResult{}, err
	}

	return ctx.Result, nil
}
