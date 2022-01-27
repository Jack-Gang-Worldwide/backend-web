package database

import (
	"database/sql"
	"fmt"
	"log"
)

func GetDB(username string, password string, address string, dbName string) *sql.DB{
	log.Printf("INFO Starting database connection")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Printf("ERROR Failed to open connection with database : %v", err)
	}
	if db.Ping(); err != nil {
		log.Printf("ERROR Failed to ping the database : %v", err);
	}
	return db
}