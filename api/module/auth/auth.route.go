package auth

import (
	"web-chat/types"
)

type AuthRoute interface {
	Route()
}
type authRoute struct {
	controller AuthController
	router     *types.Router
}

func NewAuthRoute(controller AuthController, router *types.Router) AuthRoute {
	return &authRoute{
		controller: controller,
		router:     router,
	}
}

func (route *authRoute) Route() {

	route.router.Open.POST("/register", route.controller.registerHandler)
	route.router.Open.POST("/login", route.controller.loginHandler)
	route.router.Open.POST("/forgot-password", route.controller.forgotPasswordHandler)

	route.router.Restricted.POST("/change-password", route.controller.changePasswordHandler)
}
