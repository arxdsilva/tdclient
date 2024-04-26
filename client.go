package tdclient

import (
	httpadapter "github.com/arxdsilva/tdclient/adapters/http"
	"github.com/arxdsilva/tdclient/api"
)

type Client struct {
	api.API
}

// New creates a new instance of the Client.
// It accepts optional ClientOption arguments to customize the configuration.
// Returns a pointer to the newly created Client.
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
