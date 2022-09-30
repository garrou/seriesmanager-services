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
	routes := e.Group("/api", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.POST("/seasons", s.Post)
		routes.GET("/series/:id/seasons", s.GetDistinctBySeriesId)
		routes.GET("/series/:id/seasons/:number", s.GetInfosBySeasonBySeriesId)
		routes.GET("/series/:id/seasons/viewed", s.GetDetailsSeasonsNbViewed)
		routes.GET("/seasons/continue", s.GetToContinue)
		routes.PATCH("/seasons/:id", s.Update)
		routes.DELETE("/seasons/:id", s.Delete)
	}
}

// Post user adds a season
func (s *seasonController) Post(ctx *gin.Context) {
	var seasonsDto dto.SeasonsCreateDto

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.AddSeasonsBySeries(userId, seasonsDto.SeriesId, seasonsDto)

	if seasons, ok := res.(dto.SeasonsCreateDto); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Saison(s) ajoutée(s)", seasons))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible d'ajouter les saison", nil))
	}
}

// GetDistinctBySeriesId gets series seasons by series sid
func (s *seasonController) GetDistinctBySeriesId(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seasons := s.seasonService.GetDistinctBySeriesId(userId, seriesId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", seasons))
}

// GetInfosBySeasonBySeriesId get season user infos
func (s *seasonController) GetInfosBySeasonBySeriesId(ctx *gin.Context) {
	seriesId, errId := strconv.Atoi(ctx.Param("id"))
	number, errNum := strconv.Atoi(ctx.Param("number"))

	if errId != nil || errNum != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seasonService.GetInfosBySeasonBySeriesId(userId, seriesId, number)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// GetDetailsSeasonsNbViewed get the number of each season was viewed
func (s *seasonController) GetDetailsSeasonsNbViewed(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seasonService.GetDetailsSeasonsNbViewed(userId, seriesId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// GetToContinue get user's series with unwatched seasons
func (s *seasonController) GetToContinue(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seasonService.GetToContinue(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// Update updates season viewedAt
func (s *seasonController) Update(ctx *gin.Context) {
	var seasonsDto dto.SeasonUpdateDto

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := s.seasonService.UpdateSeason(userId, seasonsDto)

	if _, ok := res.(entities.Season); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série modifiée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier la saison", nil))
	}
}

// Delete delete one user's season
func (s *seasonController) Delete(ctx *gin.Context) {
	seasonId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	isDeleted := s.seasonService.DeleteSeason(userId, seasonId)

	if isDeleted {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Saison supprimée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de supprimer la saison", nil))
	}
}
