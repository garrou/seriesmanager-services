package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type SeriesController interface {
	Routes(e *gin.Engine)
	PostSeries(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByTitle(ctx *gin.Context)
	GetInfosBySid(ctx *gin.Context)
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
		routes.GET("/titles/:title", s.GetByTitle)
		routes.GET("/:sid/infos", s.GetInfosBySid)
		routes.DELETE("/:sid", s.Delete)
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
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	seriesDto.User = fmt.Sprintf("%s", claims["id"])

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
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%s", claims["id"])
	series := s.seriesService.GetAll(userId)
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// GetByTitle returns all series with title matching
func (s *seriesController) GetByTitle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%s", claims["id"])
	series := s.seriesService.GetByTitle(userId, ctx.Param("title"))
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosBySid returns user infos series by sid
func (s *seriesController) GetInfosBySid(ctx *gin.Context) {
	infos := s.seriesService.GetInfosBySid(ctx.Param("sid"))
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}

// Delete deletes series with userId and sid
func (s *seriesController) Delete(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%s", claims["id"])
	isDeleted := s.seriesService.DeleteByUserBySid(userId, ctx.Param("sid"))

	if isDeleted {
		response := helpers.NewResponse("", nil)
		ctx.JSON(http.StatusNoContent, response)
	} else {
		response := helpers.NewResponse("Une erreur est survenue durant la suppression de la série", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}
