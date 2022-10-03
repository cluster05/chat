package websocket

import (
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
	// chatCollections := datasource.MongoDB.Database(DBMongo).Collection(CollectionChat)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("[socket][connect]")

		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("[socket][dis-connect]")
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		log.Println("[socket][connection error]", err)
	})

	server.OnEvent("/chat", "join", func(socket socketio.Conn, friendshipId string) {
		log.Println("[server][join room] ", friendshipId)
		server.JoinRoom("/chat", friendshipId, socket)
	})

	server.OnEvent("/chat", "message", func(socket socketio.Conn, personalChatDTO personal.PersonalChatDTO) {

		log.Println("[server][receive message] ", personalChatDTO)

		chat := personal.PersonalChat{
			PersonalChatId: ksuid.New().String(),
			FriendshipId:   personalChatDTO.FriendshipId,
			From:           personalChatDTO.From,
			To:             personalChatDTO.To,
			Message:        personalChatDTO.Message,
			CreatedAt:      time.Now().Unix(),
			UpdatedAt:      time.Now().Unix(),
		}

		log.Println(chat)

		// ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
		// chatCollections.InsertOne(ctx, chat)

		server.BroadcastToRoom("/chat", personalChatDTO.FriendshipId, "message", chat)
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
