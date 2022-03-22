package dto

import "github.com/Jack-Gang-Worldwide/backend-web/internal/members/entity"

type MemberResponse struct {
	Id			uint64 `json:"id"`
	RealName    string `json:"real_name"`
	Birth       int    `json:"birth"`
	Born        string `json:"born"`
	Description string `json:"description"`
}

type MembersResponse []MemberResponse

func NewMemberResponse(m entity.Member) MemberResponse {
	return MemberResponse{
		Id: m.Id,
		RealName: m.RealName,
		Birth: m.Birth,
		Born: m.Born,
		Description: m.Description,
	}
}

func NewMembersResponse(m entity.Members) *MembersResponse{
	var membersResponse MembersResponse

	for _, each := range m{
		member := NewMemberResponse(*each)
		membersResponse = append(membersResponse, member)
	}

	return &membersResponse
}