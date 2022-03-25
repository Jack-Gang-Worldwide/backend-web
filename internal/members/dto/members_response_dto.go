package dto

import (
	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/entity"
	"github.com/peteprogrammer/go-automapper"
)

type InsertMemberResponse struct {
	Id          uint64 `json:"id"`
	RealName    string `json:"real_name"`
	Birth       int    `json:"birth"`
	Born        string `json:"born"`
	Description string `json:"description"`
}

type InsertMembersResponse []InsertMemberResponse

func CreateNewMemberResponse(memberEntity entity.Member) *InsertMemberResponse{
	var insertMemberResponse InsertMemberResponse
	automapper.Map(memberEntity, &insertMemberResponse)
	return &insertMemberResponse
}