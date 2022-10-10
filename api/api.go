package api

import (
	"log"

	"github.com/cluster05/chat/api/middlewere"
	"github.com/cluster05/chat/database"
	"github.com/cluster05/chat/websocket"

	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {

	datasource, err := database.InitDatabase()
	if err != nil {
		return nil, err
	}
	log.Println("[api][InitDatabase][done]")

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middlewere.CORSMiddleware())
	log.Println("[api][GinSetup][done]")

	websocket.InitSocketIO(router, datasource)
	log.Println("[api][Websocket][done]")

	HandlerSetup(router, datasource)
	log.Println("[api][HandlerSetup][done]")
	return router, nil
}
