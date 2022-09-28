package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/services"
	"strconv"
)

// CreateSeason user adds a season
func CreateSeason(ctx *gin.Context) {
	var seasonsDto dto.SeasonsCreateDto

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.AddSeasonsBySeries(userId, seasonsDto.SeriesId, seasonsDto)

	if seasons, ok := res.(dto.SeasonsCreateDto); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Saison(s) ajoutée(s)", seasons))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible d'ajouter les saison", nil))
	}
}

// GetDistinctBySeriesId gets series seasons by series sid
func GetDistinctBySeriesId(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	seasons := services.GetDistinctBySeriesId(userId, seriesId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", seasons))
}

// GetInfosBySeasonBySeriesId get season user infos
func GetInfosBySeasonBySeriesId(ctx *gin.Context) {
	seriesId, errId := strconv.Atoi(ctx.Param("id"))
	number, errNum := strconv.Atoi(ctx.Param("number"))

	if errId != nil || errNum != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := services.GetInfosBySeasonBySeriesId(userId, seriesId, number)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// GetDetailsSeasonsNbViewed get the number of each season was viewed
func GetDetailsSeasonsNbViewed(ctx *gin.Context) {
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := services.GetDetailsSeasonsNbViewed(userId, seriesId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// GetToContinue get user's series with unwatched seasons
func GetToContinue(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	series := services.GetToContinue(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// UpdateSeason updates season viewedAt
func UpdateSeason(ctx *gin.Context) {
	var seasonsDto dto.SeasonUpdateDto

	if errDto := ctx.ShouldBind(&seasonsDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.UpdateSeason(userId, seasonsDto)

	if _, ok := res.(entities.Season); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série modifiée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier la saison", nil))
	}
}

// DeleteSeason delete one user's season
func DeleteSeason(ctx *gin.Context) {
	seasonId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	isDeleted := services.DeleteSeason(userId, seasonId)

	if isDeleted {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Saison supprimée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de supprimer la saison", nil))
	}
}
