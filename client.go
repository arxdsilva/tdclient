package tdclient

import (
	"net/http"

	"github.com/arxdsilva/tdclient/adapters"
	httpadapter "github.com/arxdsilva/tdclient/adapters/http"
)

type Client struct {
	cfg  *Config
	http adapters.Http
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

	return &Client{
		cfg:  cfg,
		http: httpadapter.New(client, cfg.url()),
	}, nil
}
