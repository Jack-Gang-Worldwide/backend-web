package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/dto"
)

type ApexLegendsService interface {
	GetAllApexMembers(ctx context.Context) (dto.MembersResponse, error)
}