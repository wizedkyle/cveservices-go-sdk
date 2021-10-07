# cveservices-go-sdk

`cvesservices-go-sdk` is the CVE Services SDK for the Go programming language.

The SDK requires a minimum version of `Go 1.16`.

## Getting Started

To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with `go get`. 
This example shows how you can use the SDK to make an API request to retrieve information about a CVE ID.

###### Add SDK Dependencies

`go get github.com/wizedkle/cveservices-go-sdk`

###### Write Code

```go
package main

import (
	"fmt"
	"github.com/wizedkyle/cveservices-go-sdk"
	"net/http"
)

func main() {
	client := cveservices_go_sdk.APIClient{
		Cfg: &cveservices_go_sdk.Configuration{
			Authentication: cveservices_go_sdk.BasicAuth{
				APIUser: "<apiUser>",
				APIKey:  "<apiKey>",
			},
			BasePath:     "<dev or prod url>",
			Organization: "<organizationShortName>",
			HTTPClient:   &http.Client{},
		},
	}
	data, response, err := client.GetCveId("<cveId>")
	if err != nil {
		fmt.Println(err)
    } else {
		if response.StatusCode == 200 {
			fmt.Println(data)
        }
    }
}
```

## Contributing

The CVE Services Go SDK uses GitHub [Issues](https://github.com/wizedkyle/cveservices-go-sdk/issues) to report and track
issues with the SDK. If you have found a bug, identified an area of improvement or want a new feature add please create or
upvote an existing issue.

You can open pull requests for the CVE Services Go SDK for fixes or additions, these pull requests will be reviewed by the maintainers before being merged.
