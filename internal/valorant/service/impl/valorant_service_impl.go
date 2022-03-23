package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/dto"
	valorantRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type valorantServiceImpl struct {
	rr valorantRepositorypkg.ValorantRepository
}

func ProvideValorantService(rr valorantRepositorypkg.ValorantRepository) *valorantServiceImpl{
	return &valorantServiceImpl{rr: rr}
}

func(vs valorantServiceImpl)GetAllValorantMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := vs.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllValorantMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0{
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}

	members, realname, err := vs.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllValorantMembers -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateMembersResponse(members, realname), nil
}