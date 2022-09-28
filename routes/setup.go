package routes

import (
	"github.com/gin-gonic/gin"
	"seriesmanager-services/controllers"
	"seriesmanager-services/middlewares"
)

func Setup(e *gin.Engine) {
	e.POST("/api/register", controllers.Register)
	e.POST("/api/login", controllers.Login)

	e.GET("/api/search/discover", middlewares.AuthorizeJwt(), controllers.Discover)
	e.GET("/api/search/names", middlewares.AuthorizeJwt(), controllers.Get)
	e.GET("/api/search/names/:name", middlewares.AuthorizeJwt(), controllers.GetApiSeriesByName)
	e.GET("/api/search/series/:sid/seasons", middlewares.AuthorizeJwt(), controllers.GetSeasonsBySid)
	e.GET("/api/search/series/:sid/seasons/:number/episodes", middlewares.AuthorizeJwt(), controllers.GetEpisodesBySidBySeason)
	e.GET("/api/search/names/:name/images", middlewares.AuthorizeJwt(), controllers.GetImagesBySeriesName)

	e.POST("/api/seasons", middlewares.AuthorizeJwt(), controllers.CreateSeason)
	e.GET("/api/series/:id/seasons", middlewares.AuthorizeJwt(), controllers.GetDistinctBySeriesId)
	e.GET("/api/series/:id/seasons/:number", middlewares.AuthorizeJwt(), controllers.GetInfosBySeasonBySeriesId)
	e.GET("/api/series/:id/seasons/viewed", middlewares.AuthorizeJwt(), controllers.GetDetailsSeasonsNbViewed)
	e.GET("/api/seasons/continue", middlewares.AuthorizeJwt(), controllers.GetToContinue)
	e.PATCH("/api/seasons/:id", middlewares.AuthorizeJwt(), controllers.UpdateSeason)
	e.DELETE("/api/seasons/:id", middlewares.AuthorizeJwt(), controllers.DeleteSeason)

	e.POST("/api/series/", middlewares.AuthorizeJwt(), controllers.CreateSeries)
	e.GET("/api/series/", middlewares.AuthorizeJwt(), controllers.GetAllSeries)
	e.GET("/api/series/names", middlewares.AuthorizeJwt(), controllers.GetSeriesByName)
	e.GET("/api/series/names/:name", middlewares.AuthorizeJwt(), controllers.GetSeriesByName)
	e.GET("/api/series/:id", middlewares.AuthorizeJwt(), controllers.GetSeriesInfosById)
	e.DELETE("/api/series/:id", middlewares.AuthorizeJwt(), controllers.DeleteSeries)
	e.PATCH("/api/series/:id/watching", middlewares.AuthorizeJwt(), controllers.UpdateWatching)

	e.GET("/api/user/", middlewares.AuthorizeJwt(), controllers.GetUser)
	e.PATCH("/api/user/profile", middlewares.AuthorizeJwt(), controllers.UpdateProfile)
	e.PATCH("/api/user/banner", middlewares.AuthorizeJwt(), controllers.UpdateBanner)
	e.PATCH("/api/user/password", middlewares.AuthorizeJwt(), controllers.UpdatePassword)

	e.GET("/api/stats/series/years", middlewares.AuthorizeJwt(), controllers.GetAddedSeriesByYears)
	e.GET("/api/stats/series/count", middlewares.AuthorizeJwt(), controllers.GetTotalSeries)
	e.GET("/api/stats/seasons/years", middlewares.AuthorizeJwt(), controllers.GetNbSeasonsByYears)
	e.GET("/api/stats/seasons/months", middlewares.AuthorizeJwt(), controllers.GetNbSeasonsByMonths)
	e.GET("/api/stats/seasons/time", middlewares.AuthorizeJwt(), controllers.GetTimeSeasonsByYears)
	e.GET("/api/stats/episodes/years", middlewares.AuthorizeJwt(), controllers.GetNbEpisodesByYears)
	e.GET("/api/stats/time", middlewares.AuthorizeJwt(), controllers.GetTotalTime)
	e.GET("/api/stats/time/month", middlewares.AuthorizeJwt(), controllers.GetTimeCurrentMonth)
}
