package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/arxdsilva/tdclient/adapters"
	"github.com/arxdsilva/tdclient/api/models"
)

var (
	_ API = &client{}
)

const (
	worldPath  = "/world/%s"
	worldsPath = "/worlds"

	characterPath = "/character/%s"

	guildPath = "/guild/%s"
)

type client struct {
	http   adapters.Http
	url    string
	logger *slog.Logger
}

// NewClient creates a new instance of the client.
// It takes an httpClient of type adapters.Http and a logger of type *slog.Logger as parameters.
// It returns a pointer to the created client.
func NewClient(httpClient adapters.Http, logger *slog.Logger) *client {
	return &client{
		http:   httpClient,
		logger: logger,
	}
}

// GetWorld retrieves information about a specific world.
// It takes a `world` parameter specifying the world to retrieve.
// It returns a pointer to `models.V4GetWorldResponse` and an error.
// If the request is successful, the response will contain the information about the world.
// If an error occurs during the request or response parsing, an error will be returned.
func (c *client) GetWorld(world string) (*models.V4GetWorldResponse, error) {
	response, status, err := c.http.DoRequest(
		http.MethodGet, fmt.Sprintf(worldPath, world), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetWorldResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetWorld error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, json.Unmarshal(response, resp)
}

// GetWorlds retrieves a list of worlds from the server.
// It returns a pointer to a V4GetWorldsResponse struct and an error, if any.
func (c *client) GetWorlds() (*models.V4GetWorldsResponse, error) {
	response, status, err := c.http.DoRequest(
		http.MethodGet, worldsPath, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetWorldsResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetWorlds error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetCharacter retrieves information about a specific character.
// It takes a `character` parameter specifying the character to retrieve.
// It returns a pointer to `models.V4GetCharacterResponse` and an error.
// If the request is successful, the response will contain the information about the character.
// If an error occurs during the request or response parsing, an error will be returned.
func (c *client) GetCharacter(character string) (*models.V4GetCharacterResponse, error) {
	response, status, err := c.http.DoRequest(
		http.MethodGet, fmt.Sprintf(characterPath, character), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetCharacterResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetCharacter error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetGuild retrieves information about a specific guild.
// It takes a `guildName` parameter specifying the guild to retrieve.
// It returns a pointer to `models.V4GetGuildResponse` and an error.
// If the request is successful, the response will contain the information about the guild.
// If an error occurs during the request or response parsing, an error will be returned.
func (c *client) GetGuild(guildName string) (*models.V4GetGuildResponse, error) {
	guildName = strings.ReplaceAll(guildName, " ", "%20")
	response, status, err := c.http.DoRequest(
		http.MethodGet, fmt.Sprintf(guildPath, guildName), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetGuildResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetGuild error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}
