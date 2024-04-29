package api

import (
	"context"

	"github.com/arxdsilva/tdclient/api/models"
)

type API interface {
	// worlds
	GetWorld(ctx context.Context, world string) (*models.V4GetWorldResponse, error)
	GetWorlds(ctx context.Context) (*models.V4GetWorldsResponse, error)

	// characters
	GetCharacter(ctx context.Context, character string) (*models.V4GetCharacterResponse, error)

	// guilds
	GetGuild(ctx context.Context, guildName string) (*models.V4GetGuildResponse, error)
	GetGuilds(ctx context.Context, world string) (*models.V4GetGuildsResponse, error)

	// bosses
	GetBoostableBosses(ctx context.Context) (*models.V4GetBoostableBossesResponse, error)

	// creatures
	GetCreatures(ctx context.Context) (*models.V4GetCreaturesResponse, error)
	GetCreatureByName(ctx context.Context, name string) (*models.V4GetCreatureByNameResponse, error)

	// fansites
	GetFansites(ctx context.Context) (*models.V4GetFansitesResponse, error)

	// highscores
	GetHighScores(ctx context.Context, world, category, vocation string, page int) (*models.V4GetHighScores, error)

	// houses
	GetHousesByWorldTown(ctx context.Context, world, town string) (*models.V4GetHousesByWorldAndTownResponse, error)
	GetHouseByWorldAndID(ctx context.Context, world, houseID string) (*models.V4GetHouseByTownIDResponse, error)
}
