package impl

import "database/sql"

type theforestRepositoryImpl struct {
	DB *sql.DB
}

func ProvideTheForestRepository(DB *sql.DB) *theforestRepositoryImpl{
	return &theforestRepositoryImpl{DB: DB}
}