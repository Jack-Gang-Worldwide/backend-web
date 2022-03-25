package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/entity"

type MemberResponse struct {
	Id          uint64 `json:"id"`
	MemberId    uint64 `json:"member_id"`
	InGameName  string `json:"ingame_name"`
	RealName    string `json:"real_name"`
	Image       string `json:"image"`
	Role        string `json:"role"`
	Nationality string `json:"nationality"`
}

type MembersResponse []MemberResponse

func CreateMemberResponse(mm entity.MobileLegendsMember, m entity.Member)MemberResponse{
	return MemberResponse{
		Id: mm.Id,
		MemberId: mm.MemberId,
		InGameName: mm.InGameName,
		RealName: m.RealName,
		Image: mm.Image,
		Role: mm.Role,
		Nationality: mm.Nationality,
	}
}

func CreateMembersResponse(mm entity.MobileLegendsMembers, m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for idx := range mm{
		members := CreateMemberResponse(*mm[idx], *m[idx])
		membersResponse = append(membersResponse, members)
	}

	return &membersResponse
}