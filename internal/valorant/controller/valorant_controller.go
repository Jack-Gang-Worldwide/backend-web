package controller

import (
	"net/http"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	valorantServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/service/api"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type ValorantController struct {
	router *mux.Router
	vs valorantServicePkg.ValorantService
}

func(vc *ValorantController)viewAllValorantMembers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := vc.vs.GetAllValorantMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(vc *ValorantController) InitializeController(){
	//TODO: Add your routes here
	vc.router.HandleFunc(global.API_PATH_GET_ALL_VALORANT_MEMBERS, vc.viewAllValorantMembers).Methods("GET")
}

func ProvideValorantController(router *mux.Router, vs valorantServicePkg.ValorantService) *ValorantController{
	return &ValorantController{router: router, vs: vs}
}