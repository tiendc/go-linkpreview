package linkpreview

import "strings"

func (ctx *ParserContext) parseTwitterMeta() {
	if !ctx.Config.ParseTwitterMeta {
		return
	}

	twitterMeta := &TwitterMeta{}
	for _, tag := range ctx.MetaTags {
		tagName := strOr(tag.Name, tag.Property)
		if !strings.HasPrefix(tagName, "twitter:") {
			continue
		}
		tagName = tagName[len("twitter:"):]
		tagContent := strOr(tag.Content, tag.Value)

		switch tagName {
		case "url":
			twitterMeta.URL = tagContent
		case "card":
			twitterMeta.Card = tagContent
		case "site":
			twitterMeta.Site = tagContent
		case "site:id":
			twitterMeta.SiteID = tagContent
		case "creator":
			twitterMeta.Creator = tagContent
		case "creator:id":
			twitterMeta.CreatorID = tagContent
		case "title":
			twitterMeta.Title = tagContent
		case "description":
			twitterMeta.Description = tagContent
		case "image":
			twitterMeta.Image = tagContent
		case "image:alt":
			twitterMeta.ImageAlt = tagContent
		default:
			twitterMeta.Others = append(twitterMeta.Others, tag)
		}
	}

	ctx.Result.TwitterMeta = twitterMeta
}
