package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/entity"
)

type TheForestRepository interface {
	GetAllMembers(ctx context.Context) (entity.TheForestMembers, entity.Members, error)
	CheckMembersExists(ctx context.Context) (int, error)
	GetListMembers(ctx context.Context)(entity.Members, error)
}