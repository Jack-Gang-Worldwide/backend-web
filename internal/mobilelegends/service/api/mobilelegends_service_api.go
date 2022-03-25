package api

import (
	"context"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/dto"
)

type MobileLegendsService interface {
	GetAllMobileLegendsMembers(ctx context.Context) (dto.MembersResponse, error)
}