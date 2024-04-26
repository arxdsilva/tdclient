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

// WithEnv sets the environment for the client configuration.
// It takes an `env` string as a parameter and returns a `ClientOption` function.
// The `ClientOption` function modifies the `Config` object by setting the `Env` field.
// If the `env` is not equal to "api", the `Env` field is set to "dev".
// Otherwise, the `Env` field is set to the value of `env`.
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

// WithTimeout sets the timeout duration for the HTTP client in the configuration.
// It returns a ClientOption function that modifies the given Config object.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Config) *Config {
		c.httpClient.Timeout = timeout
		return c
	}
}

// WithLogger sets the logger for the client.
// It takes a pointer to a slog.Logger and returns a ClientOption function.
// The ClientOption function sets the logger in the provided Config and returns the updated Config.
func WithLogger(lg *slog.Logger) ClientOption {
	return func(c *Config) *Config {
		c.logger = lg
		return c
	}
}

// WithLogLevel sets the log level for the client.
// It takes a slog.Level as input and returns a ClientOption.
// The ClientOption modifies the Config object by setting the log level.
func WithLogLevel(lv slog.Level) ClientOption {
	return func(c *Config) *Config {
		slog.SetLogLoggerLevel(lv)
		return c
	}
}

// WithHTTPClient is a ClientOption function that sets the HTTP client for the configuration.
// It takes a pointer to an http.Client and returns a function that sets the client in the Config struct.
func WithHTTPClient(cl *http.Client) ClientOption {
	return func(c *Config) *Config {
		c.httpClient = cl
		return c
	}
}
