package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/services"
)

type UserController interface {
	Routes(e *gin.Engine)
	Get(ctx *gin.Context)
	UpdateBanner(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtHelper   helpers.JwtHelper
}

func NewUserController(userService services.UserService, jwtHelper helpers.JwtHelper) UserController {
	return &userController{userService: userService, jwtHelper: jwtHelper}
}

func (u *userController) Routes(e *gin.Engine) {
	routes := e.Group("/api/user", middlewares.AuthorizeJwt(u.jwtHelper))
	{
		routes.GET("/", u.Get)
		routes.PATCH("/profile", u.UpdateProfile)
		routes.PATCH("/banner", u.UpdateBanner)
		routes.PATCH("/password", u.UpdatePassword)
	}
}

// Get gets the authenticated user
func (u *userController) Get(ctx *gin.Context) {
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Get(userId)

	if user, ok := res.(entities.User); ok {
		response := helpers.NewResponse("", dto.UserDto{
			Username: user.Username,
			Email:    user.Email,
			JoinedAt: user.JoinedAt,
			Banner:   user.Banner,
		})
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewResponse("Impossible de récupérer votre profil", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// UpdateBanner updates the banner of the authenticated user
func (u *userController) UpdateBanner(ctx *gin.Context) {
	var body struct {
		Banner string `json:"banner"`
	}
	_ = ctx.Bind(&body)
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.UpdateBanner(userId, body.Banner)

	if _, ok := res.(entities.User); ok {
		response := helpers.NewResponse("Bannière modifiée", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewResponse("Impossible de modifier la bannière", nil)
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
}

// UpdateProfile updates the authenticated user account
func (u *userController) UpdateProfile(ctx *gin.Context) {
	var userDto dto.UserUpdateProfileDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userDto.Id = u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.UpdateProfile(userDto)

	if _, ok := res.(entities.User); ok {
		response := helpers.NewResponse("Profil modifié", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewResponse("Impossible de modifier le profil", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// UpdatePassword updates authenticated user password
func (u *userController) UpdatePassword(ctx *gin.Context) {
	var userDto dto.UserUpdatePasswordDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userDto.Id = u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.UpdatePassword(userDto)

	if _, ok := res.(entities.User); ok {
		response := helpers.NewResponse("Mot de passe modifié", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewResponse("Impossible de modifier le mot de passe", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
