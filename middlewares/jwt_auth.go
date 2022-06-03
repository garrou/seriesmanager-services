package middlewares

import (
	"net/http"
	"seriesmanager-services/helpers"

	"github.com/gin-gonic/gin"
)

func AuthorizeJwt(jwtHelper helpers.JwtHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helpers.NewErrorResponse("Aucun token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := jwtHelper.ValidateToken(authHeader)

		if err != nil {
			panic(err.Error())
		}
		if !token.Valid {
			response := helpers.NewErrorResponse("Token invalide", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
