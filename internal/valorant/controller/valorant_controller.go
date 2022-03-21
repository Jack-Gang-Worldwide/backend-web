package controller

import (
	valorantServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/service/api"
	"github.com/gorilla/mux"
)

type ValorantController struct {
	router *mux.Router
	vs valorantServicePkg.ValorantService
}

func(vs *ValorantController) InitializeController(){
	//TODO: Add your routes here
}

func ProvideValorantController(router *mux.Router, vs valorantServicePkg.ValorantService) *ValorantController{
	return &ValorantController{router: router, vs: vs}
}