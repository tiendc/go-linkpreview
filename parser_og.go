package linkpreview

import "strings"

func (ctx *ParserContext) parseOGMeta() {
	if !ctx.Config.ParseOGMeta {
		return
	}

	ogMeta := &OGMeta{}
	var imageTags []*MetaTag
	var videoTags []*MetaTag

	for _, tag := range ctx.MetaTags {
		tagProp := tag.Property
		if !strings.HasPrefix(tagProp, "og:") {
			continue
		}
		tagContent := tag.Content
		parsed := true

		switch tagProp {
		case "og:url":
			ogMeta.URL = tagContent
		case "og:title":
			ogMeta.Title = tagContent
		case "og:type":
			ogMeta.Type = tagContent
		case "og:description":
			ogMeta.Description = tagContent
		case "og:site_name":
			ogMeta.SiteName = tagContent
		case "og:locale":
			ogMeta.Locale = tagContent
		default:
			parsed = false
		}
		if parsed {
			continue
		}
		if strings.HasPrefix(tagProp, "og:image") {
			imageTags = append(imageTags, tag)
		} else if strings.HasPrefix(tagProp, "og:video") {
			videoTags = append(videoTags, tag)
		}
	}

	if len(imageTags) > 0 {
		ctx.parseOGImages(ogMeta, imageTags)
	}
	if len(videoTags) > 0 {
		ctx.parseOGVideos(ogMeta, videoTags)
	}

	ctx.Result.OGMeta = ogMeta
}

func (ctx *ParserContext) parseOGImages(ogMeta *OGMeta, tags []*MetaTag) {
	var currImage *OGImage
	for _, tag := range tags {
		switch tag.Property {
		case "og:image":
			if currImage != nil {
				ogMeta.Images = append(ogMeta.Images, currImage)
			}
			currImage = &OGImage{URL: parseURL(tag.Content, ctx.Link)}
		case "og:image:url", "og:image:secure_url":
			if currImage != nil {
				currImage.URL = tag.Content
			}
		case "og:image:width":
			if currImage != nil {
				currImage.Width = parseInt(tag.Content)
			}
		case "og:image:height":
			if currImage != nil {
				currImage.Height = parseInt(tag.Content)
			}
		case "og:image:type":
			if currImage != nil {
				currImage.Type = tag.Content
			}
		case "og:image:alt":
			if currImage != nil {
				currImage.Alt = tag.Content
			}
		}
	}
	if currImage != nil {
		ogMeta.Images = append(ogMeta.Images, currImage)
	}
}

func (ctx *ParserContext) parseOGVideos(ogMeta *OGMeta, tags []*MetaTag) {
	var currVideo *OGVideo
	for _, tag := range tags {
		switch tag.Property {
		case "og:video":
			if currVideo != nil {
				ogMeta.Videos = append(ogMeta.Videos, currVideo)
			}
			currVideo = &OGVideo{URL: parseURL(tag.Content, ctx.Link)}
		case "og:video:url", "og:video:secure_url":
			if currVideo != nil {
				currVideo.URL = tag.Content
			}
		case "og:video:width":
			if currVideo != nil {
				currVideo.Width = parseInt(tag.Content)
			}
		case "og:video:height":
			if currVideo != nil {
				currVideo.Height = parseInt(tag.Content)
			}
		case "og:video:type":
			if currVideo != nil {
				currVideo.Type = tag.Content
			}
		}
	}
	if currVideo != nil {
		ogMeta.Videos = append(ogMeta.Videos, currVideo)
	}
}
