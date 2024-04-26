package tdclient

import (
	"time"
)

type ClientOption func(*Config) *Config

type Config struct {
	Env     string
	Version string
	Host    string
	timeout time.Duration
}

func defaultCfg() *Config {
	return &Config{
		Env:     "api",
		Version: "v4",
		Host:    ".tibiadata.com",
		timeout: time.Minute,
	}
}

func (c *Config) url() string {
	return "https://" + c.Env + c.Host + "/" + c.Version + "/"
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
