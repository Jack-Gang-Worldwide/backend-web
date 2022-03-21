package impl

import (
	theforestRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/repository/api"
)

type theforestServiceImpl struct {
	rr theforestRepositorypkg.TheForestRepository
}

func ProvideTheForestService(rr theforestRepositorypkg.TheForestRepository) *theforestServiceImpl{
	return &theforestServiceImpl{rr: rr}
}