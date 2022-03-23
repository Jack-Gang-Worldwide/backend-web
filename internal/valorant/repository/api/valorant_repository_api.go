package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/entity"
)

type ValorantRepository interface {
	GetListMembers(ctx context.Context)(entity.Members, error)
	GetAllMembers(ctx context.Context) (entity.ValorantMembers, entity.Members, error)
	CheckMembersExists(ctx context.Context)(int, error)
}