package api

import (
	"web-chat/api/middlewere"
	"web-chat/database"
	"web-chat/websocket"

	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {

	datasource, err := database.InitDatabase()
	if err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middlewere.CORSMiddleware())

	websocket.InitSocketIO(router, datasource)

	HandlerSetup(router, datasource)
	return router, nil
}
