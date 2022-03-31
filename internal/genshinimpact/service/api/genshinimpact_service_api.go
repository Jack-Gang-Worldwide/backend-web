package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/dto"
)

type GenshinImpactService interface {
	GetAllGenshinMembers(ctx context.Context) (dto.MembersResponse, error)
}