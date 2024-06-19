package linkpreview

import (
	"github.com/PuerkitoBio/goquery"
)

type MetaTag struct {
	Name       string            `json:"name,omitempty"`
	Property   string            `json:"property,omitempty"`
	ItemProp   string            `json:"itemprop,omitempty"`
	Content    string            `json:"content,omitempty"`
	Value      string            `json:"value,omitempty"`
	OtherAttrs map[string]string `json:"other_attrs,omitempty"`
}

type LinkTag struct {
	Rel        string            `json:"rel,omitempty"`
	Href       string            `json:"href,omitempty"`
	Sizes      string            `json:"sizes,omitempty"`
	Type       string            `json:"type,omitempty"`
	OtherAttrs map[string]string `json:"other_attrs,omitempty"`
}

type Image struct {
	URL    string `json:"url"`
	Type   string `json:"type,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Alt    string `json:"alt,omitempty"`
}

type OGImage struct {
	URL    string `json:"url"`
	Type   string `json:"type,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Alt    string `json:"alt,omitempty"`
}

type OGVideo struct {
	URL    string `json:"url"`
	Type   string `json:"type,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type OGMeta struct {
	URL         string     `json:"url,omitempty"`
	Title       string     `json:"title,omitempty"`
	Type        string     `json:"type,omitempty"`
	Description string     `json:"description,omitempty"`
	SiteName    string     `json:"site_name,omitempty"`
	Locale      string     `json:"locale,omitempty"`
	Images      []*OGImage `json:"images,omitempty"`
	Videos      []*OGVideo `json:"videos,omitempty"`
	Others      []*MetaTag `json:"others,omitempty"`
}

type TwitterMeta struct {
	URL         string     `json:"url,omitempty"`
	Card        string     `json:"card,omitempty"`
	Site        string     `json:"site,omitempty"`
	SiteID      string     `json:"site_id,omitempty"`
	Creator     string     `json:"creator,omitempty"`
	CreatorID   string     `json:"creator_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Image       string     `json:"image,omitempty"`
	ImageAlt    string     `json:"image_alt,omitempty"`
	Others      []*MetaTag `json:"others,omitempty"`
}

type ParserConfig struct {
	ParseOGMeta      bool
	ParseTwitterMeta bool
	ParseFavicons    bool

	ReturnMetaTags bool
	ReturnLinkTags bool
}

type ParserResult struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	OGMeta      *OGMeta      `json:"og_meta,omitempty"`
	TwitterMeta *TwitterMeta `json:"twitter_meta,omitempty"`
	Favicons    []*Image     `json:"favicons,omitempty"`

	MetaTags []*MetaTag `json:"meta_tags,omitempty"`
	LinkTags []*LinkTag `json:"link_tags,omitempty"`
}

type ParserContext struct {
	Link   string
	Config ParserConfig

	MetaTags []*MetaTag
	LinkTags []*LinkTag

	Result ParserResult

	Doc *goquery.Document
}

type ConfigOption func(*ParserConfig)

func ParseOGMeta(flag bool) ConfigOption {
	return func(config *ParserConfig) {
		config.ParseOGMeta = flag
	}
}

func ParseTwitterMeta(flag bool) ConfigOption {
	return func(config *ParserConfig) {
		config.ParseTwitterMeta = flag
	}
}

func ParseFavicons(flag bool) ConfigOption {
	return func(config *ParserConfig) {
		config.ParseFavicons = flag
	}
}

func ReturnMetaTags(flag bool) ConfigOption {
	return func(config *ParserConfig) {
		config.ReturnMetaTags = flag
	}
}

func ReturnLinkTags(flag bool) ConfigOption {
	return func(config *ParserConfig) {
		config.ReturnLinkTags = flag
	}
}
