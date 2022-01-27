package main

import (
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/pkg/constant"
	"github.com/Jack-Gang-Worldwide/backend-web/pkg/database"
)

func main() {
	envConstant := constant.InsertEnvConstants()
	db := database.GetDB(envConstant.DB_USERNAME,envConstant.DB_PASSWORD,envConstant.DB_ADDRESS,envConstant.DB_NAME)
	log.Println(db)
}