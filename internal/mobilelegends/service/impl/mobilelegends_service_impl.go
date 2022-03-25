package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/dto"
	mobilelegendsRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type mobilelegendsServiceImpl struct {
	rr mobilelegendsRepositorypkg.MobileLegendsRepository
}

func ProvideMobileLegendsService(rr mobilelegendsRepositorypkg.MobileLegendsRepository) *mobilelegendsServiceImpl{
	return &mobilelegendsServiceImpl{rr: rr}
}

func(ms mobilelegendsServiceImpl)GetAllMobileLegendsMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := ms.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMobileLegendsMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0{
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}

	members, realname, err := ms.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMobileLegendsMembers -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateMembersResponse(members, realname), nil
}