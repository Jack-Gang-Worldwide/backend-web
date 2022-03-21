package impl

import (
	apexlegendsRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/repository/api"
)

type apexLegendsServiceImpl struct {
	rr apexlegendsRepositorypkg.ApexLegendsRepository
}

func ProvideApexLegendsService(rr apexlegendsRepositorypkg.ApexLegendsRepository) *apexLegendsServiceImpl{
	return &apexLegendsServiceImpl{rr: rr}
}