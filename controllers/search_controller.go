package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type SearchController interface {
	Routes(e *gin.Engine)
	GetSeriesByName(ctx *gin.Context)
	GetSeriesById(ctx *gin.Context)
	GetSeasonsBySeriesId(ctx *gin.Context)
}

type searchController struct {
	searchService services.SearchService
	jwtHelper     helpers.JwtHelper
}

func NewSearchController(searchService services.SearchService, jwtHelper helpers.JwtHelper) SearchController {
	return &searchController{searchService: searchService, jwtHelper: jwtHelper}
}

func (s *searchController) Routes(e *gin.Engine) {
	routes := e.Group("/api/search", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.GET("/discover", s.Discover)
		routes.GET("/names/:name", s.GetSeriesByName)
		routes.GET("/series/:id", s.GetSeriesById)
		routes.GET("/series/:id/seasons", s.GetSeasonsBySeriesId)
	}
}

// Discover calls api and returns random series
func (s *searchController) Discover(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	series := s.searchService.Discover()
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeriesByName calls api to get series by name
func (s *searchController) GetSeriesByName(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	series := s.searchService.SearchSeriesByName(ctx.Param("name"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeriesById calls api to get series details by series id api
func (s *searchController) GetSeriesById(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	series := s.searchService.SearchSeriesById(ctx.Param("id"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeasonsBySeriesId calls api to get seasons by series id api
func (s *searchController) GetSeasonsBySeriesId(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	seasons := s.searchService.SearchSeasonsBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse(true, "", seasons)
	ctx.JSON(http.StatusOK, response)
}
