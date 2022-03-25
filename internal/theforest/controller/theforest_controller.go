package controller

import (
	"net/http"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	theforestServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/service/api"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type TheForestController struct {
	router *mux.Router
	fs theforestServicePkg.TheForestService
}

func(fc *TheForestController)viewAllTheForestMebers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := fc.fs.GetAllTheForestMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(fc *TheForestController) InitializeController(){
	//TODO: Add your routes here
	fc.router.HandleFunc(global.API_PATH_GET_ALL_THE_FOREST_MEMBERS, fc.viewAllTheForestMebers).Methods("GET")
}

func ProvideTheForestController(router *mux.Router, fs theforestServicePkg.TheForestService) *TheForestController{
	return &TheForestController{router: router, fs: fs}
}