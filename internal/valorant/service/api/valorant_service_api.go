package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/dto"
)

type ValorantService interface {
	GetAllValorantMembers(ctx context.Context) (dto.MembersResponse, error)
}