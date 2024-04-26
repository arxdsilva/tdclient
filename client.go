package tdclient

import (
	httpadapter "github.com/arxdsilva/tdclient/adapters/http"
	"github.com/arxdsilva/tdclient/api"
)

type Client struct {
	api.API
}

func New(opts ...ClientOption) *Client {
	cfg := defaultCfg()
	for _, opt := range opts {
		opt(cfg)
	}

	c := &Client{
		API: api.NewClient(
			httpadapter.New(cfg.httpClient, cfg.url()),
			cfg.logger,
		),
	}
	return c
}
