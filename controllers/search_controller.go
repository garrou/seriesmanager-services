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
	SearchSeriesByName(ctx *gin.Context)
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
		routes.GET("/names/:name", s.SearchSeriesByName)
		routes.GET("/series/:id", s.SearchSeriesById)
	}
}

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

func (s *searchController) SearchSeriesByName(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	series := s.searchService.SearchSeriesByName(ctx.Param("name"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}

func (s *searchController) SearchSeriesById(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := s.jwtHelper.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}
	series := s.searchService.SearchSeriesById(ctx.Param("id"))
	response := helpers.NewResponse(true, "", series)
	ctx.JSON(http.StatusOK, response)
}
