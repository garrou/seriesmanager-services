package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type SeasonController interface {
	Routes(e *gin.Engine)
	PostSeason(ctx *gin.Context)
	GetDistinctBySeriesId(ctx *gin.Context)
	GetInfosBySeasonBySeriesId(ctx *gin.Context)
	GetDetailsSeasonsNbViewed(ctx *gin.Context)
	PostAllSeasonsSeries(ctx *gin.Context)
	GetToContinue(ctx *gin.Context)
}

type seasonController struct {
	seasonService services.SeasonService
	jwtHelper     helpers.JwtHelper
}

func NewSeasonController(seasonService services.SeasonService, jwtHelper helpers.JwtHelper) SeasonController {
	return &seasonController{seasonService: seasonService, jwtHelper: jwtHelper}
}

func (s *seasonController) Routes(e *gin.Engine) {
	routes := e.Group("/api/seasons", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.POST("/", s.PostSeason)
		routes.POST("/series/:id/all", s.PostAllSeasonsSeries)
		routes.GET("/series/:id", s.GetDistinctBySeriesId)
		routes.GET("/:number/series/:id/infos", s.GetInfosBySeasonBySeriesId)
		routes.GET("/series/:id/viewed", s.GetDetailsSeasonsNbViewed)
		routes.GET("/continue", s.GetToContinue)
	}
}

// PostSeason user adds a season
func (s *seasonController) PostSeason(ctx *gin.Context) {
	var seasonDto dto.SeasonCreateDto
	if errDto := ctx.ShouldBind(&seasonDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := s.seasonService.AddSeason(seasonDto)

	if season, ok := res.(entities.Season); ok {
		response := helpers.NewResponse("Saison ajoutée", season)
		ctx.JSON(http.StatusCreated, response)
	} else {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// GetDistinctBySeriesId gets series seasons by series sid
func (s *seasonController) GetDistinctBySeriesId(ctx *gin.Context) {
	seasons := s.seasonService.GetDistinctBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse("", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosBySeasonBySeriesId get season user infos
func (s *seasonController) GetInfosBySeasonBySeriesId(ctx *gin.Context) {
	infos := s.seasonService.GetInfosBySeasonBySeriesId(ctx.Param("id"), ctx.Param("number"))
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}

// GetDetailsSeasonsNbViewed get the number of each season was viewed
func (s *seasonController) GetDetailsSeasonsNbViewed(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seasonService.GetDetailsSeasonsNbViewed(userId, ctx.Param("id"))
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}

// PostAllSeasonsSeries allows user to add all seasons of a series
func (s *seasonController) PostAllSeasonsSeries(ctx *gin.Context) {
	var seasonsDto dto.SeasonsCreateAllDto
	var response helpers.Response
	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		response = helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.AddAllSeasonsBySeries(userId, ctx.Param("id"), seasonsDto)

	if seasons, ok := res.(dto.SeasonsCreateAllDto); ok {
		response = helpers.NewResponse("Saisons ajoutées", seasons)
		ctx.JSON(http.StatusOK, response)
	} else {
		response = helpers.NewResponse("Erreur durant l'ajout des saisons", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}

// GetToContinue get user's series with unwatched seasons
func (s *seasonController) GetToContinue(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seasonService.GetToContinue(userId)
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}
