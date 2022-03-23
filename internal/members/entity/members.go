package entity

import "time"

type Member struct {
	Id          uint64	`db:"id"`
	RealName    string	`db:"real_name"`
	Birth       int	`db:"birth"`
	Born        string	`db:"born"`
	Description string	`db:"description"`
	CreatedAt   time.Time	`db:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at"`
}

type Members []*Member