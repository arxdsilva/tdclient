# tdclient
A golang client for tibiadata's API

![Actions on main](https://github.com/arxdsilva/tdclient/actions/workflows/test.yml/badge.svg?branch=main)
[![Coverage Status](https://coveralls.io/repos/github/arxdsilva/tdclient/badge.svg?branch=main)](https://coveralls.io/github/arxdsilva/tdclient?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/tdclient)](https://goreportcard.com/report/github.com/arxdsilva/tdclient)
[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/arxdsilva/tdclient?status.svg)](https://godoc.org/github.com/arxdsilva/tdclient)

## Disclaimer

This is a work in progress client library, it does not contain yet all the routes that TibiaData supports.

## Usage

Check `example/simple/main.go` for a working example.

```go
package main

import (
	"log/slog"

    "github.com/arxdsilva/tdclient"
)

func main() {
    client := tdclient.New()
    world, err := client.GetWorld("Premia")
    if err != nil {
        // handle error
    }

    worlds, err := client.GetWorlds()
    if err != nil {
        // handle error
    }

    fmt.Println("world: ", world)
    fmt.Println("worlds:", worlds)

    // you can also configure the client with options:
    tdclient.New(
        // change the env to query from
        tdclient.WithEnv("dev"),
        // bring your own slog logger
        tdclient.WithLogger(customSlogger),
        // set the debug level
        tdclient.WithLogLevel(slog.LevelDebug),
        // add your custom http client
        tdclient.WithHTTPClient(http.DefaultClient), 
        // customise the http client max timeout duration
        tdclient.WithTimeout(time.Hour),
    )
}
```

