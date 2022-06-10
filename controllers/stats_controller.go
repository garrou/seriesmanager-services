package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type StatsController interface {
	Routes(e *gin.Engine)
	GetNbSeriesByYears(ctx *gin.Context)
	GetTimeSeriesByYears(ctx *gin.Context)
}

type statsController struct {
	statsService services.StatsService
	jwtHelper    helpers.JwtHelper
}

func NewStatsController(statsService services.StatsService, jwtHelper helpers.JwtHelper) StatsController {
	return &statsController{statsService: statsService, jwtHelper: jwtHelper}
}

func (s *statsController) Routes(e *gin.Engine) {
	routes := e.Group("/api/stats", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.GET("/seasons/years", s.GetNbSeriesByYears)
		routes.GET("/seasons/time", s.GetTimeSeriesByYears)
	}
}
func (s *statsController) GetNbSeriesByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetNbSeasonsByYears(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetTimeSeriesByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTimeSeasonsByYears(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}
