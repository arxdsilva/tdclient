package tdclient

import (
	"log/slog"
	"time"
)

type ClientOption func(*Config) *Config

type Config struct {
	Env     string
	Version string
	Host    string
	timeout time.Duration
	logger  *slog.Logger
}

func defaultCfg() *Config {
	return &Config{
		Env:     "api",
		Version: "v4",
		Host:    ".tibiadata.com",
		timeout: time.Minute,
		logger:  slog.Default(),
	}
}

func (c *Config) url() string {
	return "https://" + c.Env + c.Host + "/" + c.Version
}

func WithEnv(env string) ClientOption {
	return func(c *Config) *Config {
		if env != "api" {
			c.Env = "dev"
			return c
		}
		c.Env = env
		return c
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Config) *Config {
		c.timeout = timeout
		return c
	}
}
