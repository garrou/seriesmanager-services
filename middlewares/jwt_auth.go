package middlewares

import (
	"net/http"
	"services-series-manager/helpers"

	"github.com/gin-gonic/gin"
)

func AuthorizeJwt(jwtHelper helpers.JwtHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helpers.NewErrorResponse("Aucun token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, _ := jwtHelper.ValidateToken(authHeader)

		if !token.Valid {
			response := helpers.NewErrorResponse("Token invalide", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
