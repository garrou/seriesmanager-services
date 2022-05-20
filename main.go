package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"services-series-manager/controllers"
	"services-series-manager/database"
	"services-series-manager/helpers"
	"services-series-manager/repositories"
	"services-series-manager/services"

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
)

func main() {

	defer database.Close(db)

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.0"})

	if err != nil {
		log.Fatal(err)
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	authController.Routes(router)
	userController.Routes(router)
	searchController.Routes(router)
	seriesController.Routes(router)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
