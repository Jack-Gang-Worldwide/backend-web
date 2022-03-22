package controller

import (
	"database/sql"

	apexlegendsControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/controller"
	apexlegendsRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/repository/impl"
	apexlegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/service/impl"

	dota2ControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/dota2/controller"
	dota2RepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/dota2/repository/impl"
	dota2ServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/dota2/service/impl"

	fortniteControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/fortnite/controller"
	fortniteRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/fortnite/repository/impl"
	fortniteServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/fortnite/service/impl"

	genshinimpactControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/controller"
	genshinimpactRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/repository/impl"
	genshinimpactServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/service/impl"

	mobilelegendsControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/controller"
	mobilelegendsRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/repository/impl"
	mobilelegendsServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/service/impl"

	theforestControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/controller"
	theforestRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/repository/impl"
	theforestServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/service/impl"

	valorantControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/controller"
	valorantRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/repository/impl"
	valorantServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/service/impl"

	memberControllerPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/controller"
	memberRepositoryPkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/repository/impl"
	memberServicePkg "github.com/Jack-Gang-Worldwide/backend-web/internal/members/service/impl"

	"github.com/SIC-Unud/sicgolib"
	"github.com/gorilla/mux"
)

func InitializeController(router *mux.Router, db *sql.DB){
	router.Use(sicgolib.ErrorHandlingMiddleware)

	webrouter := router.PathPrefix(API_WEB_PATH_ROOT).Subrouter()

	apexLegendsRepository := apexlegendsRepositoryPkg.ProvideApexLegendsRepository(db)
	apexLegendsService := apexlegendsServicePkg.ProvideApexLegendsService(apexLegendsRepository)
	apexLegendsController := apexlegendsControllerPkg.ProvideApexLegendsController(webrouter, apexLegendsService)
	apexLegendsController.InitializeController()

	dota2Repository := dota2RepositoryPkg.ProvideDota2Repository(db)
	dota2Service := dota2ServicePkg.ProvideDota2Service(dota2Repository)
	dota2Controller := dota2ControllerPkg.ProvideDota2Controller(webrouter, dota2Service)
	dota2Controller.InitializeController()

	fortniteRepository := fortniteRepositoryPkg.ProvideFortniteRepository(db)
	fortniteService := fortniteServicePkg.ProvideFortniteService(fortniteRepository)
	fortniteController := fortniteControllerPkg.ProvideFortniteController(webrouter, fortniteService)
	fortniteController.InitializeController()

	genshinimpactRepository := genshinimpactRepositoryPkg.ProvideGenshinImpactRepository(db)
	genshinimpactService := genshinimpactServicePkg.ProvideGenshinImpactService(genshinimpactRepository)
	genshinimpactController := genshinimpactControllerPkg.ProvideGenshinImpactController(webrouter, genshinimpactService)
	genshinimpactController.InitializeController()

	mobilelegendsRepository := mobilelegendsRepositoryPkg.ProvideMobileLegendsRepository(db)
	mobilelegendsService := mobilelegendsServicePkg.ProvideMobileLegendsService(mobilelegendsRepository)
	mobilelegendsController := mobilelegendsControllerPkg.ProvideMobileLegendsController(webrouter, mobilelegendsService)
	mobilelegendsController.InitializeController()

	theforestRepository := theforestRepositoryPkg.ProvideTheForestRepository(db)
	theforestService := theforestServicePkg.ProvideTheForestService(theforestRepository)
	theforestController := theforestControllerPkg.ProvideTheForestController(webrouter, theforestService)
	theforestController.InitializeController()

	valorantRepository := valorantRepositoryPkg.ProvideValorantRepository(db)
	valorantService := valorantServicePkg.ProvideValorantService(valorantRepository)
	valorantController := valorantControllerPkg.ProvideValorantController(webrouter, valorantService)
	valorantController.InitializeController()

	memberRepository := memberRepositoryPkg.ProvideMemberRepository(db)
	memberService := memberServicePkg.ProvideMembersService(memberRepository)
	memberController := memberControllerPkg.ProvideMembersController(webrouter, memberService)
	memberController.InitializeController()
}