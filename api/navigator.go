package api

import (
	"net/http"

	"github.com/cluster05/chat/api/middlewere"
	"github.com/cluster05/chat/api/module/auth"
	"github.com/cluster05/chat/api/module/chat/friend"
	"github.com/cluster05/chat/api/module/chat/personal"
	"github.com/cluster05/chat/database"
	"github.com/cluster05/chat/pkg/response"
	"github.com/cluster05/chat/types"

	"github.com/gin-gonic/gin"
)

func HandlerSetup(router *gin.Engine, datasource *database.DataSource) {

	r := router.Group("/api/v1/r", middlewere.Auth())
	o := router.Group("/api/v1/o")

	o.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.Health{
			Status:  http.StatusOK,
			Message: "ok",
			Health:  "up",
		})
	})

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
	friend.NewAuthRoute(friend.NewFrinshipController(friend.NewFriendshipService(*datasource)), router).Route()
	personal.NewPersonalChatRoute(personal.NewPersonalChatController(personal.NewPersonalChatService(*datasource)), router).Route()
}
