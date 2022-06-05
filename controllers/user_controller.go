package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/models"
	"seriesmanager-services/services"
)

type UserController interface {
	Routes(e *gin.Engine)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
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
		routes.GET("/profile", u.GetProfile)
		routes.PATCH("/profile", u.Update)
	}
}

// Get gets the authenticated user
func (u *userController) Get(ctx *gin.Context) {
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Get(userId)

	if _, ok := res.(models.User); ok {
		ctx.Status(http.StatusOK)
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

// GetProfile gets the user's profile
func (u *userController) GetProfile(ctx *gin.Context) {
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Get(userId)

	if user, ok := res.(models.User); ok {
		response := helpers.NewResponse("", dto.UserProfileDto{
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

// Update updates the authenticated user account
func (u *userController) Update(ctx *gin.Context) {
	var userDto dto.UserUpdateDto
	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userDto.Id = u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Update(userDto)

	if _, ok := res.(models.User); ok {
		response := helpers.NewResponse("Profil modifié", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.NewResponse("Impossible de modifier le profil", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
