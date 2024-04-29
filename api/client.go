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
	worldPath  string = "/world/%s"
	worldsPath string = "/worlds"

	// character
	characterPath string = "/character/%s"

	// guild
	guildPath         string = "/guilds/%s"
	guildsInWorldPath string = "/guilds/%s"

	// bosses
	boostableBossesPath string = "/boostablebosses"

	// creatures
	creaturesPath       string = "/creatures"
	creaturesByNamePath string = "/creature/%s"

	// fansites
	fansitesPath string = "/fansites"

	// highscores
	highscoresPath string = "/highscores/%s/%s/%s/%d"

	// houses
	houseIDPath string = "/house/%s/%s"
	housesPath  string = "/houses/%s/%s"
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

// GetGuilds retrieves information about a specific guild.
// It takes a `guildName` parameter specifying the guild to retrieve.
// It returns a pointer to `models.V4GetGuildResponse` and an error.
// If the request is successful, the response will contain the information about the guild.
// If an error occurs during the request or response parsing, an error will be returned.
func (c *client) GetGuilds(ctx context.Context, world string) (*models.V4GetGuildsResponse, error) {
	world = strings.ToLower(world)
	response, status, err := c.http.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf(guildsInWorldPath, world), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetGuildsResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetGuilds error",
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
		http.MethodGet, fmt.Sprintf(creaturesByNamePath, creatureName), nil, nil)
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
		http.MethodGet, fansitesPath, nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetFansitesResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Fansites.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetFansites error",
			"http_status", status,
			"api_status_message", resp.Fansites.Information.Status.Message,
			"api_status_error", resp.Fansites.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetHighScores retrieves the high scores for a specific world, category, vocation, and page.
// It takes a context, `ctx`, for cancellation and deadline propagation,
// and strings representing the world, category, vocation, and page.
// It returns a pointer to `models.V4GetHighScores` and an error.
func (c *client) GetHighScores(ctx context.Context, world, category, vocation string, page int) (*models.V4GetHighScores, error) {
	response, status, err := c.http.DoRequest(ctx, http.MethodGet,
		fmt.Sprintf(highscoresPath, world, category, vocation, page), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetHighScores{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetHighScores error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetHousesByWorldTown retrieves a list of houses from the server.
// It takes a context, `ctx`, for cancellation and deadline propagation,
// and strings representing the world and town.
// It returns a pointer to `models.V4GetHousesByWorldAndTownResponse` and an error.
func (c *client) GetHousesByWorldTown(ctx context.Context, world, town string) (*models.V4GetHousesByWorldAndTownResponse, error) {
	world = strings.ToLower(world)
	town = strings.ToLower(town)
	response, status, err := c.http.DoRequest(ctx, http.MethodGet,
		fmt.Sprintf(housesPath, world, town), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetHousesByWorldAndTownResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetHousesByWorldTown error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}

// GetHouseByWorldAndID retrieves a house by its world and ID.
// It takes a context, the world name, and the house ID as parameters.
// The function returns a pointer to a `V4GetHouseByTownIDResponse` struct and an error.
// The `V4GetHouseByTownIDResponse` struct represents the response from the API.
// If an error occurs during the request or the response status is not OK, an error is returned.
func (c *client) GetHouseByWorldAndID(ctx context.Context, world, houseID string) (*models.V4GetHouseByTownIDResponse, error) {
	world = strings.ToLower(world)
	response, status, err := c.http.DoRequest(ctx, http.MethodGet,
		fmt.Sprintf(houseIDPath, world, houseID), nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &models.V4GetHouseByTownIDResponse{}
	if err := json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if status != http.StatusOK ||
		resp.Information.Status.HTTPCode != http.StatusOK {
		c.logger.Error("GetHousesByWorldTown error",
			"http_status", status,
			"api_status_message", resp.Information.Status.Message,
			"api_status_error", resp.Information.Status.Error,
		)
		return nil, err
	}
	return resp, nil
}
