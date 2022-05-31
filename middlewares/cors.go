package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Header("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
