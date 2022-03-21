package impl

import "database/sql"

type genshinimpactRepositoryImpl struct {
	DB *sql.DB
}

func ProvideGenshinImpactRepository(DB *sql.DB) *genshinimpactRepositoryImpl{
	return &genshinimpactRepositoryImpl{DB: DB}
}