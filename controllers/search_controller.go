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
	GetSeasonsBySeriesId(ctx *gin.Context)
	GetEpisodesBySeriesIdBySeason(ctx *gin.Context)
	Get(ctx *gin.Context)
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
		routes.GET("/names", s.Get)
		routes.GET("/names/:name", s.GetSeriesByName)
		routes.GET("/series/:id/seasons", s.GetSeasonsBySeriesId)
		routes.GET("/series/:id/seasons/:number/episodes", s.GetEpisodesBySeriesIdBySeason)
	}
}

// Discover calls api and returns random series
func (s *searchController) Discover(ctx *gin.Context) {
	series := s.searchService.Discover()
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// GetSeriesByName calls api to get series by name
func (s *searchController) GetSeriesByName(ctx *gin.Context) {
	series := s.searchService.SearchSeriesByName(ctx.Param("name"))
	response := helpers.NewResponse("", series)
	ctx.JSON(http.StatusOK, response)
}

// Get returns empty body
func (s *searchController) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// GetSeasonsBySeriesId calls api to get seasons by series id
func (s *searchController) GetSeasonsBySeriesId(ctx *gin.Context) {
	seasons := s.searchService.SearchSeasonsBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse("", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetEpisodesBySeriesIdBySeason calls api to get episodes by series id and season number
func (s *searchController) GetEpisodesBySeriesIdBySeason(ctx *gin.Context) {
	episodes := s.searchService.SearchEpisodesBySeriesIdBySeason(ctx.Param("id"), ctx.Param("number"))
	response := helpers.NewResponse("", episodes)
	ctx.JSON(http.StatusOK, response)
}
