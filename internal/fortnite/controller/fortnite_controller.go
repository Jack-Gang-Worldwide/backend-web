package controller

import (
	fortniteServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/fortnite/service/api"
	"github.com/gorilla/mux"
)

type FortniteController struct {
	router *mux.Router
	fs fortniteServicePkg.FortniteService
}

func(fs *FortniteController) InitializeController(){
	//TODO: Add your routes here
}

func ProvideFortniteController(router *mux.Router, fs fortniteServicePkg.FortniteService) *FortniteController{
	return &FortniteController{router: router, fs: fs}
}