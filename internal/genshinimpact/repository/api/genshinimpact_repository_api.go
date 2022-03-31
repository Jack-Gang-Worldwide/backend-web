package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/entity"
)

type GenshinImpactRepository interface {
	GetListMembers(ctx context.Context) (entity.Members, error)
	CheckMembersExists(ctx context.Context)(int, error)
	GetAllMembers(ctx context.Context)(entity.GenshinImpactMembers, entity.Members, error)
}