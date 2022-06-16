package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
	"strconv"
)

type SeriesController interface {
	Routes(e *gin.Engine)
	PostSeries(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByName(ctx *gin.Context)
	GetInfosById(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
		routes.POST("/", s.PostSeries)
		routes.GET("/", s.GetAll)
		routes.GET("/names", s.GetByName)
		routes.GET("/names/:name", s.GetByName)
		routes.GET("/:id/infos", s.GetInfosById)
		routes.DELETE("/:id", s.Delete)
	}
}

// PostSeries adds a series to the authenticated user account
func (s *seriesController) PostSeries(ctx *gin.Context) {
	var seriesDto dto.SeriesCreateDto
	if errDto := ctx.ShouldBind(&seriesDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesDto.UserId = userId

	if s.seriesService.IsDuplicateSeries(seriesDto) {
		response := helpers.NewResponse("Vous avez déjà ajouté cette série", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		series := s.seriesService.AddSeries(seriesDto)
		response := helpers.NewResponse("Série ajoutée", series)
		ctx.JSON(http.StatusCreated, response)
	}
}

// GetAll returns all series of the authenticated user
func (s *seriesController) GetAll(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seriesService.GetAll(userId)
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// GetByName returns all series with title matching
func (s *seriesController) GetByName(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	series := s.seriesService.GetByUserIdByName(userId, ctx.Param("name"))
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosById returns series by id
func (s *seriesController) GetInfosById(ctx *gin.Context) {
	infos := s.seriesService.GetInfosBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}

// Delete deletes series with userId and sid
func (s *seriesController) Delete(ctx *gin.Context) {
	userId := s.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	isDeleted := s.seriesService.DeleteByUserIdBySeriesId(userId, seriesId)

	if isDeleted {
		response := helpers.NewResponse("", nil)
		ctx.JSON(http.StatusNoContent, response)
	} else {
		response := helpers.NewResponse("Une erreur est survenue durant la suppression de la série", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}
