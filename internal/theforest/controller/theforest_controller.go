package controller

import (
	theforestServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/service/api"
	"github.com/gorilla/mux"
)

type TheForestController struct {
	router *mux.Router
	fs theforestServicePkg.TheForestService
}

func(fs *TheForestController) InitializeController(){
	//TODO: Add your routes here
}

func ProvideTheForestController(router *mux.Router, fs theforestServicePkg.TheForestService) *TheForestController{
	return &TheForestController{router: router, fs: fs}
}