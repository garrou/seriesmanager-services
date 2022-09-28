package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"seriesmanager-services/database"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/routes"
)

func main() {

	defer database.Close(database.Db)

	gin.SetMode(os.Getenv("GIN_MODE"))

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv.Error())
	}
	database.Open()
	router := gin.Default()
	router.Use(middlewares.CORS())
	routes.Setup(router)

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err.Error())
	}
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
