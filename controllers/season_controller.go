package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type SeasonController interface {
	Routes(e *gin.Engine)
	PostSeason(ctx *gin.Context)
	GetBySeriesId(ctx *gin.Context)
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
		routes.GET("/series/:id", s.GetBySeriesId)
	}
}

// PostSeason adds a season
func (s *seasonController) PostSeason(ctx *gin.Context) {
	var seasonDto dto.SeasonCreateDto
	if errDto := ctx.ShouldBind(&seasonDto); errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	season := s.seasonService.AddSeason(seasonDto)
	response := helpers.NewResponse(true, "Saison ajoutée", season)
	ctx.JSON(http.StatusCreated, response)
}

// GetBySeriesId allows to get series seasons by series sid
func (s *seasonController) GetBySeriesId(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	seasons := s.seasonService.GetBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse(true, "", seasons)
	ctx.JSON(http.StatusOK, response)
}
