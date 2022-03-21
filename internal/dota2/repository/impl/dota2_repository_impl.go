package impl

import "database/sql"

type dota2RepositoryImpl struct {
	DB *sql.DB
}

func ProvideDota2Repository(DB *sql.DB) *dota2RepositoryImpl{
	return &dota2RepositoryImpl{DB: DB}
}