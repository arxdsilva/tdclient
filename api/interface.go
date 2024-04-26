package api

import "github.com/arxdsilva/tdclient/api/models"

type API interface {
	GetWorld(world string) (*models.V4GetWorldResponse, error)
	GetWorlds() (*models.V4GetWorldsResponse, error)
}
