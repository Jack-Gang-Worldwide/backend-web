package controller

import (
	"net/http"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/dto"
	membersServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/service/api"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type MembersController struct {
	router *mux.Router
	ms membersServicePkg.MembersService
}

func(mc *MembersController) insertNewMember(rw http.ResponseWriter, r *http.Request){
	memberRequest := new(dto.MemberRequest)
	err := memberRequest.FromJSON(r.Body)
	if err != nil {
		panic(sicgolib.NewErrorResponse(400,
			sicgolib.RESPONSE_ERROR_BUSINESS_LOGIC_MESSAGE,
			sicgolib.NewErrorResponseValue("request body", "invalid json format "+err.Error())))
	}
	res, err := mc.ms.InsertNewMember(r.Context(), *memberRequest)
	sicgolib.NewBaseResponse(201, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, res).SendResponse(&rw)
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
	mc.router.HandleFunc(global.API_PATH_INSERT_MEMBERS, mc.insertNewMember).Methods("POST")
}

func ProvideMembersController(router *mux.Router, ms membersServicePkg.MembersService) *MembersController{
	return &MembersController{router: router, ms: ms}	
}