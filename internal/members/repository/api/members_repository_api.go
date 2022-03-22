package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/entity"
)

type MembersRepository interface {
	GetAllMembers(ctx context.Context) (entity.Members, error)
	CheckMembersExists(ctx context.Context)(int, error)
}