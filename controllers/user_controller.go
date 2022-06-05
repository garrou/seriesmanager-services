package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
		routes.PATCH("/profile", u.Update)
	}
}

// Get gets the authenticated user
func (u *userController) Get(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	res := u.userService.Get(fmt.Sprintf("%s", claims["id"]))

	if _, ok := res.(models.User); ok {
		response := helpers.NewResponse("Authentifié", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.NewResponse("Non authentifié", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

// Update updates the authenticated user account
func (u *userController) Update(ctx *gin.Context) {
	var userDto dto.UserUpdateDto
	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtHelper.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userDto.Id = fmt.Sprintf("%s", claims["id"])
	res := u.userService.Update(userDto)

	if _, ok := res.(models.User); ok {
		response := helpers.NewResponse("Profil modifié", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.NewResponse("Impossible de modifier le profil", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
