package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/models"
	"seriesmanager-services/services"
)

type SeasonController interface {
	Routes(e *gin.Engine)
	PostSeason(ctx *gin.Context)
	GetDistinctBySid(ctx *gin.Context)
	GetInfosBySeasonBySeries(ctx *gin.Context)
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
		routes.GET("/series/:sid", s.GetDistinctBySid)
		routes.GET("/:number/series/:sid/infos", s.GetInfosBySeasonBySeries)
	}
}

// PostSeason user adds a season
func (s *seasonController) PostSeason(ctx *gin.Context) {
	var seasonDto dto.SeasonCreateDto
	if errDto := ctx.ShouldBind(&seasonDto); errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := s.seasonService.AddSeason(seasonDto)

	if season, ok := res.(models.Season); ok {
		response := helpers.NewResponse(true, "Saison ajoutée", season)
		ctx.JSON(http.StatusCreated, response)
	} else {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// GetDistinctBySid gets series seasons by series sid
func (s *seasonController) GetDistinctBySid(ctx *gin.Context) {
	seasons := s.seasonService.GetDistinctBySid(ctx.Param("sid"))
	response := helpers.NewResponse(true, "", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosBySeasonBySeries get season user infos
func (s *seasonController) GetInfosBySeasonBySeries(ctx *gin.Context) {
	infos := s.seasonService.GetInfosBySeasonBySeries(ctx.Param("sid"), ctx.Param("number"))
	response := helpers.NewResponse(true, "", infos)
	ctx.JSON(http.StatusOK, response)
}
