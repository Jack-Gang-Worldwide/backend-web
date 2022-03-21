package impl

import "database/sql"

type valorantRepositoryImpl struct {
	DB *sql.DB
}

func ProvideValorantRepository(DB *sql.DB) *valorantRepositoryImpl{
	return &valorantRepositoryImpl{DB: DB}
}