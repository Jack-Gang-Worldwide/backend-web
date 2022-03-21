package impl

import (
	mobilelegendsRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/repository/api"
)

type mobilelegendsServiceImpl struct {
	rr mobilelegendsRepositorypkg.MobileLegendsRepository
}

func ProvideMobileLegendsService(rr mobilelegendsRepositorypkg.MobileLegendsRepository) *mobilelegendsServiceImpl{
	return &mobilelegendsServiceImpl{rr: rr}
}