package api

import (
	"context"
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
	// world
	worldPath  = "/world/%s"
	worldsPath = "/worlds"

	// character
	characterPath = "/character/%s"

	// guild
	guildPath = "/guild/%s"

	// bosses
	boostableBossesPath = "/boostablebosses"

	// creatures
	creaturesPath     = "/creatures"
	creaturesRacePath = "/creature/%s"
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
func (c *client) GetWorld(ctx context.Context, world string) (*models.V4GetWorldResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
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
func (c *client) GetWorlds(ctx context.Context) (*models.V4GetWorldsResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
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
func (c *client) GetCharacter(ctx context.Context, character string) (*models.V4GetCharacterResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
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
func (c *client) GetGuild(ctx context.Context, guildName string) (*models.V4GetGuildResponse, error) {
	guildName = strings.ReplaceAll(guildName, " ", "%20")
	response, status, err := c.http.DoRequest(ctx,
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

// GetBoostableBosses retrieves information about a specific guild.
// It takes a `guildName` parameter specifying the guild to retrieve.
// It returns a pointer to `models.V4GetBoostableBossesResponse` and an error.
// If the request is successful, the response will contain the information about the guild.
// If an error occurs during the request or response parsing, an error will be returned.
func (c *client) GetBoostableBosses(ctx context.Context) (*models.V4GetBoostableBossesResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
		http.MethodGet, boostableBossesPath, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetBoostableBossesResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetBoostableBosses error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetCreatures retrieves a list of creatures from the server.
// It sends an HTTP GET request to the creatures endpoint and parses the response into a V4GetCreaturesResponse object.
// If the request is successful and the response status code is 200 OK, it returns the parsed response.
// Otherwise, it logs an error message and returns an error.
func (c *client) GetCreatures(ctx context.Context) (*models.V4GetCreaturesResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
		http.MethodGet, creaturesPath, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetCreaturesResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetCreatures error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetCreatureByName retrieves a creature by its name.
// It takes a context, `ctx`, for cancellation and deadline propagation,
// and a string, `name`, representing the name of the creature to retrieve.
// It returns a pointer to `models.V4GetCreatureByNameResponse` and an error.
func (c *client) GetCreatureByName(ctx context.Context, name string) (*models.V4GetCreatureByNameResponse, error) {
	creatureName := strings.ReplaceAll(strings.ToLower(name), " ", "%20")
	response, status, err := c.http.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf(creaturesRacePath, creatureName), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetCreatureByNameResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetCreatureByName error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetFansites retrieves a list of fansites from the server.
// It returns a pointer to a V4GetFansitesResponse struct and an error, if any.
func (c *client) GetFansites(ctx context.Context) (*models.V4GetFansitesResponse, error) {
	response, status, err := c.http.DoRequest(ctx,
		http.MethodGet, creaturesPath, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetFansitesResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Fansites.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetCreatures error",
			"http_status", status,
			"api_status_message", resp.Fansites.Information.Status.Message,
			"api_status_error", resp.Fansites.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}
