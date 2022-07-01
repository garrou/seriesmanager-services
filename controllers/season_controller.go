package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
	"strconv"
)

type SeasonController interface {
	Routes(e *gin.Engine)
	Post(ctx *gin.Context)
	GetDistinctBySeriesId(ctx *gin.Context)
	GetInfosBySeasonBySeriesId(ctx *gin.Context)
	GetDetailsSeasonsNbViewed(ctx *gin.Context)
	PostAllSeasons(ctx *gin.Context)
	GetToContinue(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
		routes.POST("/", s.Post)
		routes.POST("/series/:id/all", s.PostAllSeasons)
		routes.GET("/series/:id", s.GetDistinctBySeriesId)
		routes.GET("/:number/series/:id/infos", s.GetInfosBySeasonBySeriesId)
		routes.GET("/series/:id/viewed", s.GetDetailsSeasonsNbViewed)
		routes.GET("/continue", s.GetToContinue)
		routes.PATCH("/:id", s.Update)
		routes.DELETE("/:id", s.Delete)
	}
}

// Post user adds a season
func (s *seasonController) Post(ctx *gin.Context) {
	var seasonDto dto.SeasonCreateDto
	response := helpers.NewResponse("Informations invalides", nil)

	if errDto := ctx.ShouldBind(&seasonDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.AddSeason(userId, seasonDto)

	if season, ok := res.(entities.Season); ok {
		response = helpers.NewResponse("Saison ajoutée", season)
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// GetDistinctBySeriesId gets series seasons by series sid
func (s *seasonController) GetDistinctBySeriesId(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response := helpers.NewResponse("Impossible de récupérer les informations", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seasons := s.seasonService.GetDistinctBySeriesId(userId, seriesId)
	response := helpers.NewResponse("", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosBySeasonBySeriesId get season user infos
func (s *seasonController) GetInfosBySeasonBySeriesId(ctx *gin.Context) {
	seriesId, errId := strconv.Atoi(ctx.Param("id"))
	number, errNum := strconv.Atoi(ctx.Param("number"))

	if errId != nil || errNum != nil {
		response := helpers.NewResponse("Impossible de récupérer les informations", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seasonService.GetInfosBySeasonBySeriesId(userId, seriesId, number)
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)

}

// GetDetailsSeasonsNbViewed get the number of each season was viewed
func (s *seasonController) GetDetailsSeasonsNbViewed(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response := helpers.NewResponse("Impossible de récupérer les informations", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seasonService.GetDetailsSeasonsNbViewed(userId, seriesId)
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}

// PostAllSeasons allows user to add all seasons of a series
func (s *seasonController) PostAllSeasons(ctx *gin.Context) {
	var seasonsDto dto.SeasonsCreateAllDto
	response := helpers.NewResponse("Impossible d'ajouter les saisons", nil)

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.AddAllSeasonsBySeries(userId, seasonsDto.SeriesId, seasonsDto)

	if seasons, ok := res.(dto.SeasonsCreateAllDto); ok {
		response = helpers.NewResponse("Saisons ajoutées", seasons)
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// GetToContinue get user's series with unwatched seasons
func (s *seasonController) GetToContinue(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seasonService.GetToContinue(userId)
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// Update updates season viewedAt
func (s *seasonController) Update(ctx *gin.Context) {
	var seasonsDto dto.SeasonUpdateDto
	response := helpers.NewResponse("Impossible de modifier la série", nil)

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.UpdateSeason(userId, seasonsDto)

	if _, ok := res.(entities.Season); ok {
		response = helpers.NewResponse("Série modifiée", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// Delete delete one user's season
func (s *seasonController) Delete(ctx *gin.Context) {
	response := helpers.NewResponse("Impossible de supprimer la série", nil)
	seasonId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	isDeleted := s.seasonService.DeleteSeason(userId, seasonId)

	if isDeleted {
		response = helpers.NewResponse("Saison supprimée", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}
