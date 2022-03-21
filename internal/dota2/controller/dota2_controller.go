package controller

import (
	dota2ServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/dota2/service/api"
	"github.com/gorilla/mux"
)

type Dota2Controller struct {
	router *mux.Router
	ds dota2ServicePkg.Dota2Service
}

func (ds *Dota2Controller) InitializeController() {
	//TODO: Add your routes here
}

func ProvideDota2Controller(router *mux.Router, ds dota2ServicePkg.Dota2Service) *Dota2Controller{
	return &Dota2Controller{router: router, ds: ds}
}