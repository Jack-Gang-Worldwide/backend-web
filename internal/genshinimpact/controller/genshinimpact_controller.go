package controller

import (
	"net/http"

	genshinimpactServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/service/api"
	"github.com/Jack-Gang-Worldwide/backend-web/internal/global"
	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

type GenshinImpactController struct {
	router *mux.Router
	gs genshinimpactServicePkg.GenshinImpactService
}

func(gc GenshinImpactController)viewAllGenshinMembers(rw http.ResponseWriter, r *http.Request){
	allMembers, err := gc.gs.GetAllGenshinMembers(r.Context())
	if err != nil {
		panic(sicgolib.NewErrorResponse(
			400,
			sicgolib.RESPONSE_ERROR_RUNTIME_MESSAGE,
			sicgolib.NewErrorResponseValue("request error", err.Error())))
	}
	sicgolib.NewBaseResponse(200, sicgolib.RESPONSE_SUCCESS_MESSAGE, nil, allMembers).SendResponse(&rw)
}

func(gc *GenshinImpactController) InitializeController(){
	//TODO: Add your routes here
	gc.router.HandleFunc(global.API_PATH_GET_ALL_GENSHIN_MEMBERS, gc.viewAllGenshinMembers).Methods("GET")
}

func ProvideGenshinImpactController(router *mux.Router, gs genshinimpactServicePkg.GenshinImpactService) *GenshinImpactController{
	return &GenshinImpactController{router: router, gs: gs}
}