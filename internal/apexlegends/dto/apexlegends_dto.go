package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/entity"

type MemberResponse struct {
	Id          uint64 `json:"id"`
	MemberId    uint64 `json:"member_id"`
	InGameName  string `json:"ingame_name"`
	RealName    string `json:"real_name"`
	Image       string `json:"image"`
	Nationality string `json:"nationality"`
}

type MembersResponse []MemberResponse

func CreateMemberResponse(ma entity.ApexLegendsMember, m entity.Member)MemberResponse{
	return MemberResponse{
		Id: ma.Id,
		MemberId: ma.MemberId,
		InGameName: ma.InGameName,
		RealName: m.RealName,
		Image: ma.Image,
		Nationality: ma.Nationality,
	}
}

func CreateMembersResponse(ma entity.ApexLegendsMembers, m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for idx := range ma {
		membersApex := CreateMemberResponse(*ma[idx], *m[idx])
		membersResponse = append(membersResponse, membersApex)
	}
	return &membersResponse
}