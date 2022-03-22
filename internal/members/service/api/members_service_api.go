package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/dto"
)

type MembersService interface {
	GetAllMembers(ctx context.Context) (dto.MembersResponse, error)
}