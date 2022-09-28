package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/helpers"
	"seriesmanager-services/services"
)

func GetNbSeasonsByYears(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetNbSeasonsByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetNbEpisodesByYears(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetEpisodesByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetTimeSeasonsByYears(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetTimeSeasonsByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetTotalSeries(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetTotalSeries(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetTotalTime(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetTotalTime(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetTimeCurrentMonth(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetTimeCurrentMonth(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetAddedSeriesByYears(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetAddedSeriesByYears(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}

func GetNbSeasonsByMonths(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	stats := services.GetNbSeasonsByMonths(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", stats))
}
