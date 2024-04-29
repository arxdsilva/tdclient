package api

import (
	"context"

	"github.com/arxdsilva/tdclient/api/models"
)

type API interface {
	GetWorld(ctx context.Context, world string) (*models.V4GetWorldResponse, error)
	GetWorlds(ctx context.Context) (*models.V4GetWorldsResponse, error)
	GetCharacter(ctx context.Context, character string) (*models.V4GetCharacterResponse, error)
	GetGuild(ctx context.Context, guildName string) (*models.V4GetGuildResponse, error)
	GetGuilds(ctx context.Context, world string) (*models.V4GetGuildsResponse, error)
	GetBoostableBosses(ctx context.Context) (*models.V4GetBoostableBossesResponse, error)
	GetCreatures(ctx context.Context) (*models.V4GetCreaturesResponse, error)
	GetCreatureByName(ctx context.Context, name string) (*models.V4GetCreatureByNameResponse, error)
	GetFansites(ctx context.Context) (*models.V4GetFansitesResponse, error)
}
