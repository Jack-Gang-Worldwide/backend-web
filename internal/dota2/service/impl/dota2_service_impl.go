package impl

import (
	dota2RepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/dota2/repository/api"
)

type dota2ServiceImpl struct {
	rr dota2RepositoryPkg.Dota2Repository
}

func ProvideDota2Service(rr dota2RepositoryPkg.Dota2Repository) *dota2ServiceImpl{
	return &dota2ServiceImpl{rr: rr}
}