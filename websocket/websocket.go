package websocket

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func InitSocketIO(router *gin.Engine) *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
	})

	server.OnEvent("/chat", "join", func(socket socketio.Conn, roomId string) {
		server.JoinRoom("/chat", roomId, socket)
	})

	server.OnEvent("/chat", "message", func(socket socketio.Conn, roomId, message, username string) {
		if roomId == "" && username == "" {
			return
		}
		server.BroadcastToRoom("/chat", roomId, "message", message, username)
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		log.Println("[socket][connection error]", err)
	})

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalln("[socket][connection error]", err)
		}
	}()

	return server
}
