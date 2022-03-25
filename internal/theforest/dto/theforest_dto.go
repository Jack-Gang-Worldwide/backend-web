package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/entity"

type MemberResponse struct {
	Id          uint64 `json:"id"`
	MemberId    uint64 `json:"member_id"`
	InGameName  string `json:"ingame_name"`
	RealName    string `json:"real_name"`
	Image       string `json:"image"`
	Nationality string `json:"nationality"`
}

type MembersResponse []MemberResponse

func CreateMemberResponse(mt entity.TheForestMember, m entity.Member)MemberResponse{
	return MemberResponse{
		Id: mt.Id,
		MemberId: mt.MemberId,
		InGameName: mt.InGameName,
		RealName: m.RealName,
		Image: mt.Image,
		Nationality: mt.Nationality,
	}
}

func CreateMembersResponse(mt entity.TheForestMembers, m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for idx := range mt{
		membersTheForest := CreateMemberResponse(*mt[idx], *m[idx])
		membersResponse = append(membersResponse, membersTheForest)
	}

	return &membersResponse
}