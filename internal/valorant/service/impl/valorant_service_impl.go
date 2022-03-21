package impl

import (
	valorantRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/repository/api"
)

type valorantServiceImpl struct {
	rr valorantRepositorypkg.ValorantRepository
}

func ProvideValorantService(rr valorantRepositorypkg.ValorantRepository) *valorantServiceImpl{
	return &valorantServiceImpl{rr: rr}
}