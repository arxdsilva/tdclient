package tdclient

import (
	"net/http"

	httpadapter "github.com/arxdsilva/tdclient/adapters/http"
	"github.com/arxdsilva/tdclient/api"
)

type Client struct {
	cfg *Config
	api.API
}

func New(opts ...ClientOption) (*Client, error) {
	cfg := defaultCfg()
	for _, opt := range opts {
		opt(cfg)
	}

	client := http.DefaultClient
	if cfg.timeout != 0 {
		client.Timeout = cfg.timeout
	}
	c := &Client{
		cfg: cfg,
		API: api.NewClient(
			httpadapter.New(client, cfg.url())),
	}
	return c, nil
}
