package controller

import (
	"net/http"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	mobilelegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/service/api"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type MobileLegendsController struct {
	router *mux.Router
	ms mobilelegendsServicePkg.MobileLegendsService
}

func(mc *MobileLegendsController)viewAllMobileLegendsMembers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := mc.ms.GetAllMobileLegendsMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}

	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(mc *MobileLegendsController) InitializeController(){
	//TODO: Add your routes here
	mc.router.HandleFunc(global.API_PATH_GET_ALL_MOBILE_LEGENDS_MEMBERS, mc.viewAllMobileLegendsMembers).Methods("GET")
}

func ProvideMobileLegendsController(router *mux.Router, ms mobilelegendsServicePkg.MobileLegendsService) *MobileLegendsController{
	return &MobileLegendsController{router: router, ms: ms}
}