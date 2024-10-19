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

## Security Guidelines and Best Practices
## Added by VMHUNG
When using this library, it is important to follow security best practices to ensure the safety and integrity of your application. Here are some guidelines:

1. **Input Validation and Sanitization**: Always validate and sanitize input URLs before processing them. This library includes basic validation, but you should implement additional checks as needed for your specific use case.

2. **Security Headers**: The library sets some security headers for HTTP requests, such as `X-Content-Type-Options`, `X-Frame-Options`, and `X-XSS-Protection`. Ensure that these headers are appropriate for your application and consider adding more if necessary.

3. **Error Handling**: Properly handle errors to avoid exposing sensitive information. This library includes basic error handling, but you should implement additional error handling as needed.

4. **Dependencies**: Keep your dependencies up to date to avoid known vulnerabilities. Regularly check for updates and apply them promptly.

5. **Configuration**: Review and configure the library options to suit your security requirements. Disable features that you do not need to minimize the attack surface.

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
