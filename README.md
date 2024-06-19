[![Go Version][gover-img]][gover] [![GoDoc][doc-img]][doc] [![GoReport][rpt-img]][rpt]

# Link preview and site info scraping library for Go

## Installation

```shell
go get github.com/tiendc/go-linkpreview
```

## Usage

### Simple one

```go
    result, err := linkpreview.Parse(url)

    // Output:
    // {
    //   "title": <title>,
    //   "description": <description>,
    //   "og_meta": {<open graph metadata>},
    //   "twitter_meta": {<twitter metadata>},
    // }
```

### Parse with configuration

```go
    // Get everything
    result, err := linkpreview.Parse(url,
        linkpreview.ParseOGMeta(true),        // default: true
        linkpreview.ParseTwitterMeta(true),   // default: true
        linkpreview.ParseFavicons(true),      // default: false
        linkpreview.ReturnMetaTags(true),     // default: false
        linkpreview.ReturnLinkTags(true)      // default: false
    )

    // Output:
    // {
    //   "title": <title>,
    //   "description": <description>,
    //   "og_meta": {<open graph metadata>},
    //   "twitter_meta": {<twitter metadata>},
    //   "favicons": [],
    //   "meta_tags": [],
    //   "link_tags": [],
    // }
```

## Contributing

- You are welcome to make pull requests for new functions and bug fixes.

## Authors

- Dao Cong Tien ([tiendc](https://github.com/tiendc))

## License

- [MIT License](LICENSE)

[doc-img]: https://pkg.go.dev/badge/github.com/tiendc/go-linkpreview
[doc]: https://pkg.go.dev/github.com/tiendc/go-linkpreview
[gover-img]: https://img.shields.io/badge/Go-%3E%3D%201.18-blue
[gover]: https://img.shields.io/badge/Go-%3E%3D%201.18-blue
[rpt-img]: https://goreportcard.com/badge/github.com/tiendc/go-linkpreview
[rpt]: https://goreportcard.com/report/github.com/tiendc/go-linkpreview
