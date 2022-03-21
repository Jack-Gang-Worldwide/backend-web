package impl

import "database/sql"

type apexLegendsRepositoryImpl struct {
	DB *sql.DB
}

func ProvideApexLegendsRepository(DB *sql.DB) *apexLegendsRepositoryImpl{
	return &apexLegendsRepositoryImpl{DB: DB}
}