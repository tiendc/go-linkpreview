package linkpreview

import "strings"

func (ctx *ParserContext) parseFavicons() {
	if !ctx.Config.ParseFavicons {
		return
	}

	for _, tag := range ctx.LinkTags {
		if image := ctx.parseMaybeFavicon(tag); image != nil {
			ctx.Result.Favicons = append(ctx.Result.Favicons, image)
		}
	}
}

func (ctx *ParserContext) parseMaybeFavicon(tag *LinkTag) *Image {
	if !strings.Contains(tag.Rel, "icon") || tag.Href == "" {
		return nil
	}

	image := &Image{
		URL:  parseURL(tag.Href, ctx.Link),
		Type: tag.Type,
	}

	if tag.Sizes != "" {
		wh := strings.Split(tag.Sizes, "x")
		if len(wh) == 2 {
			image.Width = parseInt(wh[0])
			image.Height = parseInt(wh[1])
		}
	}

	return image
}
