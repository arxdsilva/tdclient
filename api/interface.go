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
}
