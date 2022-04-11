package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/entity"
)

type ApexLegendsRepository interface {
	CheckMembersExists(ctx context.Context) (int, error)
	GetAllMembers(ctx context.Context)(entity.ApexLegendsMembers, entity.Members, error)
}