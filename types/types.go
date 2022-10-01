package types

import "github.com/gin-gonic/gin"

type Router struct {
	Open       *gin.RouterGroup
	Restricted *gin.RouterGroup
}
