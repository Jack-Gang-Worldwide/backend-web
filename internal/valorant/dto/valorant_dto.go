package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/entity"

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

func CreateMemberResponse(mv entity.ValorantMember, m entity.Member)MemberResponse{
	return MemberResponse{
		Id: mv.Id,
		MemberId: mv.MemberId,
		InGameName: mv.InGameName,
		RealName: m.RealName,
		Image: mv.Image,
		Role: mv.Role,
		Nationality: mv.Nationality,
	}
}

func CreateMembersResponse(mv entity.ValorantMembers, m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for idx := range mv{
		membersValorant := CreateMemberResponse(*mv[idx], *m[idx])
		membersResponse = append(membersResponse, membersValorant)
	}

	return &membersResponse
}