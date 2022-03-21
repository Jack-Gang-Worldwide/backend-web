package impl

import (
	fortniteRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/fortnite/repository/api"
)

type fortniteServiceImpl struct {
	rr fortniteRepositorypkg.FortniteRepository
}

func ProvideFortniteService(rr fortniteRepositorypkg.FortniteRepository) *fortniteServiceImpl{
	return &fortniteServiceImpl{rr: rr}
}