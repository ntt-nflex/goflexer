# goflexer

![Go Flexer logo](docs/static_files/gopher_flexer.png "Go Flexer")

Goflexer is a package used for writing nFlex modules as [GoLang plugins](https://golang.org/pkg/plugin/)

## Example
```go
package main

import (
        "fmt"

        "github.com/ntt-nflex/goflexer"
)

// Test is a test handler
func Test(context goflexer.Context, event goflexer.Event) goflexer.Result {
        fmt.Println("This is a log message")
        return map[string]interface{}{
                "foo": "bar",
        }
}
```

## Build
You must create a `plugin.go` file with you code. Then build it as a Go plugin and zip it.
```sh
$ go get "github.com/ntt-nflex/goflexer"
$ go build -buildmode=plugin plugin.go
$ zip module.zip plugin.so
```

## Test
Create a ~/.flexer.yaml file if you don't already have one

```json
regions:
  default:
    cmp_api_key: *****
    cmp_api_secret: ***********
    cmp_url: https://localhost/cmp/basic/api
verify_ssl: false
```

```sh
goflexer --handler TestMethod --event '{"foo": "bar"}'
```

## Upload
Once you have the zip file, you can upload it to nFlex. The easiest way to do that is with [flexer](https://github.com/ntt-nflex/flexer)
