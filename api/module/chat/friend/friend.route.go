package friend

import (
	"github.com/cluster05/chat/types"
)

type AuthRoute interface {
	Route()
}
type authRoute struct {
	controller FriendshipController
	router     *types.Router
}

func NewAuthRoute(controller FriendshipController, router *types.Router) AuthRoute {
	return &authRoute{
		controller: controller,
		router:     router,
	}
}

func (route *authRoute) Route() {

	route.router.Restricted.POST("/friendship/create", route.controller.createFriendshipHandler)
	route.router.Restricted.POST("/friendship/get", route.controller.getFriendshipHandler)
	route.router.Restricted.POST("/friendship/delete", route.controller.deleteFriendshipHandler)
	route.router.Restricted.POST("/friendship/search", route.controller.searchFriendshipHandler)
}
