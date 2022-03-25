package entity

import "time"

type MobileLegendsMember struct {
	Id          uint64		`db:"id"`
	MemberId    uint64		`db:"member_id"`
	InGameName  string		`db:"ingame_name"`
	Image       string		`db:"image"`
	Role        string		`db:"role"`
	Nationality string		`db:"nationality"`
	CreatedAt   time.Time	`db:"created_at"`
	UpdatedAt   time.Time	`db:"updated_at"`
}

type Member struct {
	Id          uint64		`db:"id"`
	RealName    string		`db:"real_name"`
	Birth       int			`db:"birth"`
	Born        string		`db:"born"`
	Description string		`db:"description"`
	CreatedAt   time.Time	`db:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at"`
}

type Members []*Member
type MobileLegendsMembers []*MobileLegendsMember