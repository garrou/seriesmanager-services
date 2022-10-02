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

type SeriesController interface {
	Routes(e *gin.Engine)
	Post(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByName(ctx *gin.Context)
	GetInfosById(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdateWatching(ctx *gin.Context)
}

type seriesController struct {
	seriesService services.SeriesService
	jwtHelper     helpers.JwtHelper
}

func NewSeriesController(seriesService services.SeriesService, jwtHelper helpers.JwtHelper) SeriesController {
	return &seriesController{seriesService: seriesService, jwtHelper: jwtHelper}
}

func (s *seriesController) Routes(e *gin.Engine) {
	routes := e.Group("/api/series", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.POST("/", s.Post)
		routes.GET("/", s.GetAll)
		routes.GET("/names", s.GetByName)
		routes.GET("/names/:name", s.GetByName)
		routes.GET("/:id", s.GetInfosById)
		routes.DELETE("/:id", s.Delete)
		routes.PATCH("/:id/watching", s.UpdateWatching)
	}
}

// Post adds a series to the authenticated user account
func (s *seriesController) Post(ctx *gin.Context) {
	var seriesDto dto.SeriesCreateDto

	if errDto := ctx.ShouldBind(&seriesDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesDto.UserId = userId

	if s.seriesService.IsDuplicateSeries(seriesDto) {
		ctx.AbortWithStatusJSON(http.StatusConflict, helpers.NewResponse("Vous avez déjà ajouté cette série", nil))
	} else {
		series := s.seriesService.AddSeries(seriesDto)
		ctx.JSON(http.StatusCreated, helpers.NewResponse("Série ajoutée", series))
	}
}

// GetAll returns all series of the authenticated user
func (s *seriesController) GetAll(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	query, ok := ctx.GetQuery("page")

	if !ok {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse("Données invalides", nil))
		return
	}
	page, err := strconv.Atoi(query)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse("Données invalides", nil))
		return
	}
	series := s.seriesService.GetAll(userId, page)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// GetByName returns all series with title matching
func (s *seriesController) GetByName(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seriesService.GetByUserIdByName(userId, ctx.Param("name"))
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// GetInfosById returns series by id
func (s *seriesController) GetInfosById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := s.seriesService.GetInfosBySeriesId(userId, id)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// Delete deletes series with userId and id
func (s *seriesController) Delete(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	isDeleted := s.seriesService.DeleteByUserIdBySeriesId(userId, seriesId)

	if isDeleted {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série supprimée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de supprimer la série", nil))
	}
}

// UpdateWatching updates field IsWatching to avoid api call when get series to continue
func (s *seriesController) UpdateWatching(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	res := s.seriesService.UpdateWatching(userId, seriesId)

	if _, ok := res.(entities.Series); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série modifiée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier la série", nil))
	}
}
