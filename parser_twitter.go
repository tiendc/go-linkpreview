package linkpreview

import "strings"

func (ctx *ParserContext) parseTwitterMeta() {
	if !ctx.Config.ParseTwitterMeta {
		return
	}

	twitterMeta := &TwitterMeta{}
	for _, tag := range ctx.MetaTags {
		tagName := tag.Name
		if !strings.HasPrefix(tagName, "twitter:") {
			tagName = tag.Property
			if !strings.HasPrefix(tagName, "twitter:") {
				continue
			}
		}

		switch tagName {
		case "twitter:card":
			twitterMeta.Card = tag.Content
		case "twitter:site":
			twitterMeta.Site = tag.Content
		case "twitter:site:id":
			twitterMeta.SiteID = tag.Content
		case "twitter:creator":
			twitterMeta.Creator = tag.Content
		case "twitter:creator:id":
			twitterMeta.CreatorID = tag.Content
		case "twitter:title":
			twitterMeta.Title = tag.Content
		case "twitter:description":
			twitterMeta.Description = tag.Content
		case "twitter:image":
			twitterMeta.Image = tag.Content
		case "twitter:image:alt":
			twitterMeta.ImageAlt = tag.Content
		}
	}

	ctx.Result.TwitterMeta = twitterMeta
}
