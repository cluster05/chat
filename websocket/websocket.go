package websocket

import (
	"context"
	"log"
	"time"
	"web-chat/api/module/chat/personal"
	"web-chat/database"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/segmentio/ksuid"
)

var (
	DBMongo        = "chat"
	CollectionChat = "chat"
)

func InitSocketIO(router *gin.Engine, datasource *database.DataSource) *socketio.Server {

	server := socketio.NewServer(nil)
	chatCollections := datasource.MongoDB.Database(DBMongo).Collection(CollectionChat)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("[socket][connect]")
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("[socket][disconnect]")
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		log.Println("[socket][connection error]", err)
	})

	server.OnEvent("/chat", "join", func(socket socketio.Conn, authId string) {
		log.Println("[server][join room] ", authId)
		server.JoinRoom("/chat", authId, socket)
	})

	server.OnEvent("/chat", "message", func(socket socketio.Conn, payload map[string]string) {
		chat := personal.PersonalChat{
			PersonalChatId: ksuid.New().String(),
			FriendshipId:   payload["friendshipId"],
			From:           payload["from"],
			To:             payload["to"],
			Message:        payload["message"],
			CreatedAt:      time.Now().Unix(),
			UpdatedAt:      time.Now().Unix(),
			IsDeleted:      false,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err := chatCollections.InsertOne(ctx, chat)
		if err != nil {
			log.Println("[socket][chat saving error]", err)
		}

		if err == nil {
			to := payload["to"]
			from := payload["from"]

			server.BroadcastToRoom("/chat", to, "message", chat)
			server.BroadcastToRoom("/chat", from, "message", chat)
		}

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
