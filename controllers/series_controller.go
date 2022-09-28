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

// CreateSeries adds a series to the authenticated user account
func CreateSeries(ctx *gin.Context) {
	var seriesDto dto.SeriesCreateDto

	if errDto := ctx.ShouldBind(&seriesDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesDto.UserId = userId

	if services.IsDuplicateSeries(seriesDto) {
		ctx.AbortWithStatusJSON(http.StatusConflict, helpers.NewResponse("Vous avez déjà ajouté cette série", nil))
	} else {
		series := services.AddSeries(seriesDto)
		ctx.JSON(http.StatusCreated, helpers.NewResponse("Série ajoutée", series))
	}
}

// GetAllSeries returns all series of the authenticated user
func GetAllSeries(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	series := services.GetAllSeries(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// GetSeriesByName returns all series with title matching
func GetSeriesByName(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	series := services.GetByUserIdByName(userId, ctx.Param("name"))
	ctx.JSON(http.StatusOK, helpers.NewResponse("", series))
}

// GetSeriesInfosById returns series by id
func GetSeriesInfosById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	infos := services.GetInfosBySeriesId(userId, id)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", infos))
}

// DeleteSeries deletes series with userId and id
func DeleteSeries(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	isDeleted := services.DeleteByUserIdBySeriesId(userId, seriesId)

	if isDeleted {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série supprimée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de supprimer la série", nil))
	}
}

// UpdateWatching updates field IsWatching to avoid api call when get series to continue
func UpdateWatching(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	seriesId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	res := services.UpdateWatching(userId, seriesId)

	if _, ok := res.(entities.Series); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Série modifiée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier la série", nil))
	}
}
