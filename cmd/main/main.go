package main

import (
	"database/sql"
	"os"
	"strings"

	"github.com/Jack-Gang-Worldwide/backend-web/pkg/controller"
	"github.com/SIC-Unud/sicgolib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getEnvVariableValues() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")

	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")

	envVariables["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")

	return envVariables
}

func initializeDatabase(envVariables map[string]string) *sql.DB {
	return sicgolib.GetDatabase(
		envVariables["DB_ADDRESS"],
		envVariables["DB_USERNAME"],
		envVariables["DB_PASSWORD"],
		envVariables["DB_NAME"],
	)
}

func initializeGlobalRouter(envVariables map[string]string) *mux.Router {
	r := mux.NewRouter()

	arrayWhitelistedUrls := strings.Split(envVariables["WHITELISTED_URLS"], ",")

	whitelistedUrls := make(map[string]bool)

	for _, v := range arrayWhitelistedUrls {
		whitelistedUrls[v] = true
	}

	r.Use(sicgolib.ContentTypeMiddleware)
	r.Use(sicgolib.CorsMiddleware(whitelistedUrls))
	return r
}

func main() {
	godotenv.Load()
	envVariables := getEnvVariableValues()

	db := initializeDatabase(envVariables)

	r := initializeGlobalRouter(envVariables)

	controller.InitializeController(r, db)
	
	s := sicgolib.ProvideServer(envVariables["SERVER_ADDRESS"], r)
	s.ListenAndServe()
}