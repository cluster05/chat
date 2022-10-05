package personal

import (
	"web-chat/types"
)

type PersonalChatRoute interface {
	Route()
}
type personalChatRoute struct {
	controller PersonalChatController
	router     *types.Router
}

func NewPersonalChatRoute(controller PersonalChatController, router *types.Router) PersonalChatRoute {
	return &personalChatRoute{
		controller: controller,
		router:     router,
	}
}

func (route *personalChatRoute) Route() {

	route.router.Restricted.POST("/chat/personal", route.controller.getChatHandler)
}
