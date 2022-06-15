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
	GetNbSeasonsByYears(ctx *gin.Context)
	GetTimeSeasonsByYears(ctx *gin.Context)
	GetNbEpisodesByYears(ctx *gin.Context)
	GetTotalSeries(ctx *gin.Context)
	GetTotalTime(ctx *gin.Context)
	GetTimeCurrentWeek(ctx *gin.Context)
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
		routes.GET("/seasons/years", s.GetNbSeasonsByYears)
		routes.GET("/seasons/time", s.GetTimeSeasonsByYears)
		routes.GET("/episodes/years", s.GetNbEpisodesByYears)
		routes.GET("/series/count", s.GetTotalSeries)
		routes.GET("/time", s.GetTotalTime)
		routes.GET("/time/week", s.GetTimeCurrentWeek)
	}
}
func (s *statsController) GetNbSeasonsByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetNbSeasonsByYears(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetNbEpisodesByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetEpisodesByYears(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetTimeSeasonsByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTimeSeasonsByYears(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetTotalSeries(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTotalSeries(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetTotalTime(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTotalTime(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}

func (s *statsController) GetTimeCurrentWeek(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetCurrentTimeWeek(userId)
	response := helpers.NewResponse("", stats)
	ctx.JSON(http.StatusOK, response)
}
