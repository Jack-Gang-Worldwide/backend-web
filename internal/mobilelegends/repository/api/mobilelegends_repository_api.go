package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/entity"
)

type MobileLegendsRepository interface {
	GetListMembers(ctx context.Context) (entity.Members, error)
	CheckMembersExists(ctx context.Context)(int, error)
	GetAllMembers(ctx context.Context)(entity.MobileLegendsMembers, entity.Members, error)
}