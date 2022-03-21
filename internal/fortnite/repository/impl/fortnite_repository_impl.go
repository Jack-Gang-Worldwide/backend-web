package impl

import "database/sql"

type fortniteRepositoryImpl struct {
	DB *sql.DB
}

func ProvideFortniteRepository(DB *sql.DB) *fortniteRepositoryImpl{
	return &fortniteRepositoryImpl{DB: DB}
}