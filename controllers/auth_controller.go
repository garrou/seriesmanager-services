package controllers

import (
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/services"

	"github.com/gin-gonic/gin"
)

// Register creates user
func Register(ctx *gin.Context) {
	var userDto dto.UserCreateDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userDto.TrimSpace()

	if !userDto.IsValid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	if services.IsDuplicateEmail(userDto.Email) {
		ctx.AbortWithStatusJSON(http.StatusConflict, helpers.NewResponse("Un email est déjà associé à ce compte", nil))
	} else {
		services.Register(userDto)
		ctx.JSON(http.StatusCreated, helpers.NewResponse("Compte créé", nil))
	}
}

// Login authenticate user
func Login(ctx *gin.Context) {
	var userDto dto.UserLoginDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	res := services.Login(userDto.Email, userDto.Password)

	if user, ok := res.(entities.User); ok {
		token := helpers.GenerateToken(user.ID)
		ctx.JSON(http.StatusOK, helpers.NewResponse("", token))
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse("Email ou mot de passe incorrect(s)", nil))
	}
}
