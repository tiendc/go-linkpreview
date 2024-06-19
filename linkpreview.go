package linkpreview

import (
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Parse(link string, options ...ConfigOption) (ParserResult, error) {
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return ParserResult{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ParserResult{}, err
	}
	defer response.Body.Close()

	return ParseFromReader(link, response.Body, options...)
}

func ParseFromReader(link string, data io.Reader, options ...ConfigOption) (ParserResult, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return ParserResult{}, err
	}

	ctx := ParserContext{
		Link: link,
		Config: ParserConfig{
			ParseOGMeta:      true,
			ParseTwitterMeta: true,
			ParseFavicons:    false,
			ReturnMetaTags:   false,
			ReturnLinkTags:   false,
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
