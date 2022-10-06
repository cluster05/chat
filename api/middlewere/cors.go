package middlewere

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
