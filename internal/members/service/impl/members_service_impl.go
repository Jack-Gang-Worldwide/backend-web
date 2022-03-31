package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/dto"
	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/entity"
	membersRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type membersServiceImpl struct {
	rr membersRepositorypkg.MembersRepository
}

func ProvideMembersService(rr membersRepositorypkg.MembersRepository)*membersServiceImpl{
	return &membersServiceImpl{rr: rr}
}

func(ms membersServiceImpl)InsertNewMember(ctx context.Context, mr dto.MemberRequest) (uint, error) {
	id, err := ms.rr.InsertNewMember(ctx, entity.Member(mr.CreateMemberEntity()))
	if err != nil {
		panic(sicgolib.NewErrorResponse(500, sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("internal", "server error")))
	}
	return id, nil
}

func(ms membersServiceImpl)GetAllMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := ms.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0{
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}

	members, err := ms.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}

	return *dto.NewMembersResponse(members), nil
}