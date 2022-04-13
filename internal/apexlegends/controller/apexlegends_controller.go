package controller

import (
	"net/http"

	apexlegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/service/api"
	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type ApexLegendsController struct {
	router *mux.Router
	as apexlegendsServicePkg.ApexLegendsService
}

func(ac *ApexLegendsController)viewAllApexMembers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := ac.as.GetAllApexMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
		400,
		sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
		sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(ac *ApexLegendsController) InitializeController() {
	//TODO: Add your routes here
	ac.router.HandleFunc(global.API_PATH_GET_ALL_APEX_MEMBERS, ac.viewAllApexMembers).Methods("GET")
}

func ProvideApexLegendsController(router *mux.Router, as apexlegendsServicePkg.ApexLegendsService) *ApexLegendsController{
	return &ApexLegendsController{router: router, as: as}
}