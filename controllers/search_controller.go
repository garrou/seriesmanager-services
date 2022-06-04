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
	GetEpisodesBySeriesIdBySeason(ctx *gin.Context)
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
		routes.GET("/series/:id/seasons/:number/episodes", s.GetEpisodesBySeriesIdBySeason)
	}
}

// Discover calls api and returns random series
func (s *searchController) Discover(ctx *gin.Context) {
	series := s.searchService.Discover()
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeriesByName calls api to get series by name
func (s *searchController) GetSeriesByName(ctx *gin.Context) {
	series := s.searchService.SearchSeriesByName(ctx.Param("name"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeriesById calls api to get series details by series id
func (s *searchController) GetSeriesById(ctx *gin.Context) {
	series := s.searchService.SearchSeriesById(ctx.Param("id"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeasonsBySeriesId calls api to get seasons by series id
func (s *searchController) GetSeasonsBySeriesId(ctx *gin.Context) {
	seasons := s.searchService.SearchSeasonsBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse(true, "", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetEpisodesBySeriesIdBySeason calls api to get episodes by series id and season number
func (s *searchController) GetEpisodesBySeriesIdBySeason(ctx *gin.Context) {
	episodes := s.searchService.SearchEpisodesBySeriesIdBySeason(ctx.Param("id"), ctx.Param("number"))
	response := helpers.NewResponse(true, "", episodes)
	ctx.JSON(http.StatusOK, response)
}
