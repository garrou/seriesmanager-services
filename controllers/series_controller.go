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
	}
}

// PostSeries adds a series to the authenticated user account
func (s *seriesController) PostSeries(ctx *gin.Context) {
	var seriesDto dto.SeriesCreateDto
	if errDto := ctx.ShouldBind(&seriesDto); errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	seriesDto.User = fmt.Sprintf("%s", claims["id"])

	if s.seriesService.IsDuplicateSeries(seriesDto) {
		response := helpers.NewErrorResponse("Vous avez déjà ajouté cette série", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		series := s.seriesService.AddSeries(seriesDto)
		response := helpers.NewResponse(true, "Série ajoutée", series)
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
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetByTitle returns all series with title matching
func (s *seriesController) GetByTitle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := s.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%s", claims["id"])
	series := s.seriesService.GetByTitle(userId, ctx.Param("title"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}
