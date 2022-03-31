package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/dto"
	genshinimpactRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type genshinImpactServiceImpl struct {
	rr genshinimpactRepositorypkg.GenshinImpactRepository
}

func ProvideGenshinImpactService(rr genshinimpactRepositorypkg.GenshinImpactRepository) *genshinImpactServiceImpl{
	return &genshinImpactServiceImpl{rr: rr}
}

func(gs genshinImpactServiceImpl)GetAllGenshinMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := gs.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllGenshinMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0 {
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}
	
	members, realName, err := gs.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateMembersResponse(members, realName), nil
}
