package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"services-series-manager/dto"
	"services-series-manager/helpers"
	"services-series-manager/middlewares"
	"services-series-manager/services"
)

type SeriesController interface {
	Routes(e *gin.Engine)
	AddSeries(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type seriesController struct {
	seriesService services.SeriesService
	jwtHelper     helpers.JwtHelper
}

func NewSeriesController(seriesService services.SeriesService, jwtHelper helpers.JwtHelper) SeriesController {
	return &seriesController{seriesService: seriesService, jwtHelper: jwtHelper}
}

func (s *seriesController) Routes(e *gin.Engine) {
	routes := e.Group("/series", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.POST("/", s.AddSeries)
		routes.GET("/", s.GetAll)
	}
}

func (s *seriesController) AddSeries(ctx *gin.Context) {
	var seriesDto dto.SeriesCreateDto
	errDto := ctx.ShouldBind(&seriesDto)

	if errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	seriesDto.FkUser = fmt.Sprintf("%s", claims["id"])
	s.seriesService.AddSeries(seriesDto)
	response := helpers.NewResponse(true, "Série ajoutée", nil)
	ctx.JSON(http.StatusCreated, response)
}

func (s *seriesController) GetAll(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%s", claims["id"])
	res := s.seriesService.GetAll(userId)
	response := helpers.NewResponse(true, "", res)
	ctx.JSON(http.StatusOK, response)
}
