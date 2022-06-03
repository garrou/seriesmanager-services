package controllers

import (
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/models"
	"seriesmanager-services/services"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Routes(e *gin.Engine)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtHelper   helpers.JwtHelper
}

func NewAuthController(userService services.AuthService, jwtHelper helpers.JwtHelper) AuthController {
	return &authController{
		authService: userService,
		jwtHelper:   jwtHelper,
	}
}

func (a *authController) Routes(e *gin.Engine) {
	routes := e.Group("/api")
	{
		routes.POST("/register", a.Register)
		routes.POST("/login", a.Login)
	}
}

// Register creates user
func (a *authController) Register(ctx *gin.Context) {
	var userDto dto.UserCreateDto
	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if userDto.Password != userDto.Confirm {
		response := helpers.NewErrorResponse("Le mot de passe et la confirmation du mot de passe sont différents", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if a.authService.IsDuplicateEmail(userDto.Email) {
		response := helpers.NewErrorResponse("Un email est déjà associé à un compte", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		a.authService.Register(userDto)
		response := helpers.NewResponse(true, "Compte créé", nil)
		ctx.JSON(http.StatusCreated, response)
	}
}

// Login authenticate user
func (a *authController) Login(ctx *gin.Context) {
	var userDto dto.UserDto
	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := a.authService.Login(userDto.Email, userDto.Password)

	if user, ok := res.(models.User); ok {
		token := a.jwtHelper.GenerateToken(user.Id)
		response := helpers.NewResponse(true, "OK", token)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewErrorResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
