package controller

import (
	"net/http"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	membersServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/service/api"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type MembersController struct {
	router *mux.Router
	ms membersServicePkg.MembersService
}

func(mc *MembersController) viewAllMembers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := mc.ms.GetAllMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(mc *MembersController)InitializeController(){
	//TODO: Add your routes here
	mc.router.HandleFunc(global.API_PATH_GET_ALL_MEMBERS, mc.viewAllMembers).Methods("GET")
}

func ProvideMembersController(router *mux.Router, ms membersServicePkg.MembersService) *MembersController{
	return &MembersController{router: router, ms: ms}	
}