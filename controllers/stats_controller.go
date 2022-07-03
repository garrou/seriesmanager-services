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
	GetTimeCurrentMonth(ctx *gin.Context)
	GetAddedSeriesByYears(ctx *gin.Context)
	GetNbSeasonsByMonths(ctx *gin.Context)
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
		routes.GET("/series/years", s.GetAddedSeriesByYears)
		routes.GET("/series/count", s.GetTotalSeries)
		routes.GET("/seasons/years", s.GetNbSeasonsByYears)
		routes.GET("/seasons/months", s.GetNbSeasonsByMonths)
		routes.GET("/seasons/time", s.GetTimeSeasonsByYears)
		routes.GET("/episodes/years", s.GetNbEpisodesByYears)
		routes.GET("/time", s.GetTotalTime)
		routes.GET("/time/month", s.GetTimeCurrentMonth)
	}
}
func (s *statsController) GetNbSeasonsByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetNbSeasonsByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetNbEpisodesByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetEpisodesByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetTimeSeasonsByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTimeSeasonsByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetTotalSeries(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTotalSeries(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetTotalTime(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTotalTime(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetTimeCurrentMonth(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetTimeCurrentMonth(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetAddedSeriesByYears(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetAddedSeriesByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func (s *statsController) GetNbSeasonsByMonths(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := s.statsService.GetNbSeasonsByMonths(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}
