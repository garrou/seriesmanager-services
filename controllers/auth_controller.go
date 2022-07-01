package controllers

import (
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
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
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userDto.TrimSpace()

	if !userDto.IsValid() {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if a.authService.IsDuplicateEmail(userDto.Email) {
		response := helpers.NewResponse("Un email est déjà associé à ce compte", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		a.authService.Register(userDto)
		response := helpers.NewResponse("Compte créé", nil)
		ctx.JSON(http.StatusCreated, response)
	}
}

// Login authenticate user
func (a *authController) Login(ctx *gin.Context) {
	var userDto dto.UserLoginDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := a.authService.Login(userDto.Email, userDto.Password)

	if user, ok := res.(entities.User); ok {
		token := a.jwtHelper.GenerateToken(user.ID)
		response := helpers.NewResponse("", token)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
