package impl

import "database/sql"

type mobilelegendsRepositoryImpl struct {
	DB *sql.DB
}

func ProvideMobileLegendsRepository(DB *sql.DB) *mobilelegendsRepositoryImpl{
	return &mobilelegendsRepositoryImpl{DB: DB}
}