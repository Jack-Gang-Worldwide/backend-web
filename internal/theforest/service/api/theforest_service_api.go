package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/dto"
)

type TheForestService interface {
	GetAllTheForestMembers(ctx context.Context) (dto.MembersResponse, error)
}