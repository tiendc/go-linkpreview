package linkpreview

func (ctx *ParserContext) Parse() error {
	err := ctx.readAllTags()
	if err != nil {
		return err
	}

	ctx.parseBasicTags()

	if ctx.Config.ParseOGMeta {
		ctx.parseOGMeta()
	}
	if ctx.Config.ParseTwitterMeta {
		ctx.parseTwitterMeta()
	}
	if ctx.Config.ParseFavicons {
		ctx.parseFavicons()
	}

	if ctx.Config.ReturnMetaTags {
		ctx.Result.MetaTags = ctx.MetaTags
	}
	if ctx.Config.ReturnLinkTags {
		ctx.Result.LinkTags = ctx.LinkTags
	}

	return nil
}

func (ctx *ParserContext) readAllTags() error {
	// Read all <meta> tags
	metaNodes := ctx.Doc.Find("html > head > meta")
	for _, node := range metaNodes.Nodes {
		metaTag := &MetaTag{}
		for _, attr := range node.Attr {
			switch attr.Key {
			case "name":
				metaTag.Name = attr.Val
			case "property":
				metaTag.Property = attr.Val
			case "itemprop":
				metaTag.ItemProp = attr.Val
			case "content":
				metaTag.Content = attr.Val
			default:
				if metaTag.OtherAttrs == nil {
					metaTag.OtherAttrs = map[string]string{}
				}
				metaTag.OtherAttrs[attr.Key] = attr.Val
			}
		}
		ctx.MetaTags = append(ctx.MetaTags, metaTag)
	}

	// Read all <link> tags
	linkNodes := ctx.Doc.Find("html > head > link")
	for _, node := range linkNodes.Nodes {
		linkTag := &LinkTag{}
		for _, attr := range node.Attr {
			switch attr.Key {
			case "rel":
				linkTag.Rel = attr.Val
			case "href":
				linkTag.Href = attr.Val
			case "sizes":
				linkTag.Sizes = attr.Val
			case "type":
				linkTag.Type = attr.Val
			default:
				if linkTag.OtherAttrs == nil {
					linkTag.OtherAttrs = map[string]string{}
				}
				linkTag.OtherAttrs[attr.Key] = attr.Val
			}
		}
		ctx.LinkTags = append(ctx.LinkTags, linkTag)
	}

	return nil
}

func (ctx *ParserContext) parseBasicTags() {
	// Parse title
	titleNode := ctx.Doc.Find("html > head > title")
	ctx.Result.Title = titleNode.Text()

	// Parse description
	for _, tag := range ctx.MetaTags {
		if tag.Name == "description" {
			ctx.Result.Description = tag.Content
			break
		}
	}
}
