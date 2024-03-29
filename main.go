package main

import (
	"fmt"
	"log"
	"os"
	"seriesmanager-services/controllers"
	"seriesmanager-services/database"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
	"seriesmanager-services/services"

	"github.com/gin-gonic/gin"
)

var (
	db        = database.Open()
	jwtHelper = helpers.NewJwtHelper()

	userRepository = repositories.NewUserRepository(db)
	authService    = services.NewAuthService(userRepository)
	authController = controllers.NewAuthController(authService, jwtHelper)

	userService    = services.NewUserService(userRepository)
	userController = controllers.NewUserController(userService, jwtHelper)

	searchService    = services.NewSearchService()
	searchController = controllers.NewSearchController(searchService, jwtHelper)

	seriesRepository = repositories.NewSeriesRepository(db)
	seriesService    = services.NewSeriesService(seriesRepository)
	seriesController = controllers.NewSeriesController(seriesService, jwtHelper)

	seasonRepository = repositories.NewSeasonRepository(db)
	seasonService    = services.NewSeasonService(seasonRepository, seriesRepository)
	seasonController = controllers.NewSeasonController(seasonService, jwtHelper)

	statsRepository = repositories.NewStatsRepository(db)
	statsService    = services.NewStatsService(statsRepository)
	statsController = controllers.NewStatsController(statsService, jwtHelper)
)

func main() {

	defer database.Close(db)

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	authController.Routes(router)
	userController.Routes(router)
	searchController.Routes(router)
	seriesController.Routes(router)
	seasonController.Routes(router)
	statsController.Routes(router)

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err.Error())
	}
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
