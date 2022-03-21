package controller

import (
	mobilelegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/service/api"
	"github.com/gorilla/mux"
)

type MobileLegendsController struct {
	router *mux.Router
	ms mobilelegendsServicePkg.MobileLegendsService
}

func(ms *MobileLegendsController) InitializeController(){
	//TODO: Add your routes here
}

func ProvideMobileLegendsController(router *mux.Router, ms mobilelegendsServicePkg.MobileLegendsService) *MobileLegendsController{
	return &MobileLegendsController{router: router, ms: ms}
}