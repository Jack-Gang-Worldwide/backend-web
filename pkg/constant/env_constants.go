package constant

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConstants struct {
	DB_ADDRESS  string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

// InsertEnvConstants is to get the enviroments that the developer used
// so the database connection will open dynamically
func InsertEnvConstants() *EnvConstants {
	godotenv.Load()
	envConstants := &EnvConstants{}
	envConstants.DB_ADDRESS = os.Getenv("DB_ADDRESS")
	envConstants.DB_USERNAME = os.Getenv("DB_USERNAME")
	envConstants.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	envConstants.DB_NAME = os.Getenv("DB_NAME")

	log.Println(envConstants.DB_ADDRESS)
	return envConstants
}