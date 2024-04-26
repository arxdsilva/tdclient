package api

import (
	"github.com/arxdsilva/tdclient/adapters"
	"github.com/arxdsilva/tdclient/api/models"
)

type client struct {
	http adapters.Http
}

func NewClient(httpClient adapters.Http) *client {
	return &client{http: httpClient}
}

func (c *client) GetWorld(world string) (models.V4GetWorldResponse, error) {
	return models.V4GetWorldResponse{}, nil
}

func (c *client) GetWorlds() (models.V4GetWorldsResponse, error) {
	return models.V4GetWorldsResponse{}, nil
}
