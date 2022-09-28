package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
	"seriesmanager-services/helpers"
	"seriesmanager-services/services"
)

// Discover calls api and returns random series
func Discover(ctx *gin.Context) {
	series := services.Discover()
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// GetApiSeriesByName calls api to get series by name
func GetApiSeriesByName(ctx *gin.Context) {
	series := services.SearchSeriesByName(ctx.Param("name"))
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// Get returns empty body
func Get(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

// GetSeasonsBySid calls api to get seasons by series id
func GetSeasonsBySid(ctx *gin.Context) {
	sid, err := strconv.Atoi(ctx.Param("sid"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	seasons := services.SearchSeasonsBySid(sid)
	response := helpers.NewResponse("", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetEpisodesBySidBySeason calls api to get episodes by series id and season number
func GetEpisodesBySidBySeason(ctx *gin.Context) {
	sid, errId := strconv.Atoi(ctx.Param("sid"))
	number, errNum := strconv.Atoi(ctx.Param("number"))

	if errId != nil || errNum != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	episodes := services.SearchEpisodesBySidBySeason(sid, number)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", episodes))
}

// GetImagesBySeriesName calls api to get series image with his name
func GetImagesBySeriesName(ctx *gin.Context) {
	images := services.SearchImagesBySeriesName(ctx.Param("name"))
	ctx.JSON(http.StatusOK, helpers.NewResponse("", images))
}
