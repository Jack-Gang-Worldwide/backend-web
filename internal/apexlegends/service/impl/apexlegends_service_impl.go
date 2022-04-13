package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/dto"
	apexlegendsRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type apexLegendsServiceImpl struct {
	rr apexlegendsRepositorypkg.ApexLegendsRepository
}

func ProvideApexLegendsService(rr apexlegendsRepositorypkg.ApexLegendsRepository) *apexLegendsServiceImpl{
	return &apexLegendsServiceImpl{rr: rr}
}

func (as apexLegendsServiceImpl)GetAllApexMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := as.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllApexMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0 {
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}
	members, realName, err := as.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllApexMembers -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateMembersResponse(members, realName), nil
}