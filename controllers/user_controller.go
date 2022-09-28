package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/services"
)

// GetUser gets the authenticated user
func GetUser(ctx *gin.Context) {
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.GetUser(userId)

	if user, ok := res.(entities.User); ok {
		response := helpers.NewResponse("", dto.UserDto{
			Username: user.Username,
			Email:    user.Email,
			JoinedAt: user.JoinedAt,
			Banner:   user.Banner,
		})
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de récupérer votre profil", nil))
	}
}

// UpdateBanner updates the banner of the authenticated user
func UpdateBanner(ctx *gin.Context) {
	var body struct {
		Banner string `json:"banner"`
	}
	_ = ctx.Bind(&body)
	userId := helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.UpdateBanner(userId, body.Banner)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Bannière modifiée", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, helpers.NewResponse("Impossible de modifier la bannière", nil))
	}
}

// UpdateProfile updates the authenticated user account
func UpdateProfile(ctx *gin.Context) {
	var userDto dto.UserUpdateProfileDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userDto.Id = helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.UpdateProfile(userDto)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Profil modifié", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier le profil", nil))
	}
}

// UpdatePassword updates authenticated user password
func UpdatePassword(ctx *gin.Context) {
	var userDto dto.UserUpdatePasswordDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userDto.Id = helpers.ExtractUserId(ctx.GetHeader("Authorization"))
	res := services.UpdatePassword(userDto)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Mot de passe modifié", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse("Impossible de modifier le mot de passe", nil))
	}
}
