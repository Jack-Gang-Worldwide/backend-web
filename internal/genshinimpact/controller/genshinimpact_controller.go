package controller

import (
	genshinimpactServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/service/api"
	"github.com/gorilla/mux"
)

type GenshinImpactController struct {
	router *mux.Router
	gs genshinimpactServicePkg.GenshinImpactService
}

func(gs *GenshinImpactController) InitializeController(){
	//TODO: Add your routes here
}

func ProvideGenshinImpactController(router *mux.Router, gs genshinimpactServicePkg.GenshinImpactService) *GenshinImpactController{
	return &GenshinImpactController{router: router, gs: gs}
}