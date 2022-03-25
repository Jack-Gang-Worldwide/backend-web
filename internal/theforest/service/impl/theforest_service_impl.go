package impl

import (
	"context"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/dto"
	theforestRepositorypkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/repository/api"
	"github.com/SIC-Unud/sicgolib"
)

type theforestServiceImpl struct {
	rr theforestRepositorypkg.TheForestRepository
}

func ProvideTheForestService(rr theforestRepositorypkg.TheForestRepository) *theforestServiceImpl{
	return &theforestServiceImpl{rr: rr}
}

func(fs theforestServiceImpl)GetAllTheForestMembers(ctx context.Context)(dto.MembersResponse, error){
	membersCount, err := fs.rr.CheckMembersExists(ctx)
	if err != nil {
		log.Printf("ERROR GetAllTheForestMembers -> error: %v\n", err)
		return nil, err
	}
	if membersCount == 0{
		panic(sicgolib.NewErrorResponse(494, sicgolib.RESPONSE_ERROR_DATA_NOT_EXISTS_MESSAGE,
			sicgolib.NewErrorResponseValue("members", "does not exists")))
	}

	members, realname, err := fs.rr.GetAllMembers(ctx)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	return *dto.CreateMembersResponse(members, realname), nil
}