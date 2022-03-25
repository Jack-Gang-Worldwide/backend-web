package dto

import (
	"encoding/json"
	"io"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/entity"
	"github.com/peteprogrammer/go-automapper"
)

type MemberRequest struct {
	RealName    string `json:"real_name"`
	Birth       int    `json:"birth"`
	Born        string `json:"born"`
	Description string `json:"description"`
}

func (mr *MemberRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(mr)
}

func(mr MemberRequest) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(mr)
}

func(mr MemberRequest) CreateMemberEntity() entity.Member{
	var memberEntity entity.Member
	automapper.MapLoose(mr, &memberEntity)

	return memberEntity
}