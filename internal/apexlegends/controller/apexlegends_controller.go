package controller

import (
	apexlegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/service/api"
	"github.com/gorilla/mux"
)

type ApexLegendsController struct {
	router *mux.Router
	pl apexlegendsServicePkg.ApexLegendsService
}

func(pl *ApexLegendsController) InitializeController() {
	//TODO: Add your routes here
}

func ProvideApexLegendsController(router *mux.Router, pl apexlegendsServicePkg.ApexLegendsService) *ApexLegendsController{
	return &ApexLegendsController{router: router, pl: pl}
}