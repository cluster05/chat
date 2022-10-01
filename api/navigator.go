package api

import (
	"net/http"
	"web-chat/api/middlewere"
	"web-chat/api/module/auth"
	"web-chat/database"
	"web-chat/pkg/response"
	"web-chat/types"

	"github.com/gin-gonic/gin"
)

func HandlerSetup(router *gin.Engine, datasource *database.DataSource) {

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.Health{
			Status:  http.StatusOK,
			Message: "ok",
			Health:  "up",
		})
	})

	r := router.Group("/api/v1/r", middlewere.Auth())
	o := router.Group("/api/v1/o")

	route := &types.Router{
		Open:       o,
		Restricted: r,
	}

	routing(route, datasource)

	router.NoRoute(func(ctx *gin.Context) {
		response.NotFound(ctx, "route not found")
	})

}

func routing(router *types.Router, datasource *database.DataSource) {
	auth.NewAuthRoute(auth.NewAuthController(auth.NewAuthService(*datasource)), router).Route()
}
