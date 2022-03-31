package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/entity"

type MemberResponse struct {
	Id          uint64 `json:"id"`
	MemberId    uint64 `json:"member_id"`
	InGameName  string `json:"ingame_name"`
	RealName    string `json:"real_name"`
	Image       string `json:"image"`
	Nationality string `json:"nationality"`
}

type MembersResponse []MemberResponse

func CreateMemberResponse(gi entity.GenshinImpactMember, m entity.Member) MemberResponse {
	return MemberResponse{
		Id:          gi.Id,
		MemberId:    gi.MemberId,
		InGameName:  gi.InGameName,
		RealName:    m.RealName,
		Image:       gi.Image,
		Nationality: gi.Nationality,
	}
}

func CreateMembersResponse(mv entity.GenshinImpactMembers, m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for idx := range mv{
		membersGenshin := CreateMemberResponse(*mv[idx], *m[idx])
		membersResponse = append(membersResponse, membersGenshin)
	}

	return &membersResponse
}