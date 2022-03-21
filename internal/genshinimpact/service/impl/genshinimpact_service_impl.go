package impl

import (
	genshinimpactRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/repository/api"
)

type genshinImpactServiceImpl struct {
	rr genshinimpactRepositorypkg.GenshinImpactRepository
}

func ProvideGenshinImpactService(rr genshinimpactRepositorypkg.GenshinImpactRepository) *genshinImpactServiceImpl{
	return &genshinImpactServiceImpl{rr: rr}
}
