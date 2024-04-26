package tdclient

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type ClientOption func(*Config) *Config

type Config struct {
	Env     string
	Version string
	Host    string

	httpClient *http.Client
	logger     *slog.Logger
}

func defaultCfg() *Config {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelError,
			})).With("app", "tdclient")

	dftclient := http.DefaultClient
	dftclient.Timeout = time.Minute

	return &Config{
		Env:        "api",
		Version:    "v4",
		Host:       ".tibiadata.com",
		httpClient: dftclient,
		logger:     logger,
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
		c.httpClient.Timeout = timeout
		return c
	}
}

func WithLogger(lg *slog.Logger) ClientOption {
	return func(c *Config) *Config {
		c.logger = lg
		return c
	}
}

func WithLogLevel(lv slog.Level) ClientOption {
	return func(c *Config) *Config {
		slog.SetLogLoggerLevel(lv)
		return c
	}
}

func WithHTTPClient(cl *http.Client) ClientOption {
	return func(c *Config) *Config {
		c.httpClient = cl
		return c
	}
}
