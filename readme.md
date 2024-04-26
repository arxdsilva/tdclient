# tdclient
A golang client for tibiadata's API

![Actions on main](https://github.com/arxdsilva/tdclient/actions/workflows/test.yml/badge.svg?branch=main)
[![Coverage Status](https://coveralls.io/repos/github/arxdsilva/tdclient/badge.svg?branch=main)](https://coveralls.io/github/arxdsilva/tdclient?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/tdclient)](https://goreportcard.com/report/github.com/arxdsilva/tdclient)
[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/arxdsilva/tdclient?status.svg)](https://godoc.org/github.com/arxdsilva/tdclient)

## Usage

```go
// todo
    import (
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
    }
```

