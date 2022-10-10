package personal

import (
	"context"

	"github.com/cluster05/chat/database"
	"github.com/cluster05/chat/pkg/mongoquerybuilder"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBMongo        = "chat"
	CollectionChat = "chat"
)

type PersonalChatService interface {
	createChat(context.Context, PersonalChat) error
	getChat(context.Context, string) ([]PersonalChat, error)
}

type personalChatService struct {
	Mongo *mongo.Client
}

func NewPersonalChatService(datasouce database.DataSource) PersonalChatService {
	return &personalChatService{
		Mongo: datasouce.MongoDB,
	}
}

func (pcs *personalChatService) createChat(ctx context.Context, chat PersonalChat) error {
	chatCollections := pcs.Mongo.Database(DBMongo).Collection(CollectionChat)

	_, err := chatCollections.InsertOne(ctx, chat)
	if err != nil {
		return err
	}
	return nil
}

func (pcs *personalChatService) getChat(ctx context.Context, friendshipId string) ([]PersonalChat, error) {

	chatCollections := pcs.Mongo.Database(DBMongo).Collection(CollectionChat)

	query := mongoquerybuilder.Query{
		Query: `[
			{
			  "$match": {
				"friendshipId": "~1",
				"isDeleted": false
			  }
			},
			{
			  "$sort": {
				"createdAt": 1
			  }
			}
		  ]`,
		Params: mongoquerybuilder.Params{friendshipId},
	}

	pipeline, _ := query.QueryBuilder()

	cursor, err := chatCollections.Aggregate(ctx, pipeline.Query)
	if err != nil {
		return []PersonalChat{}, err
	}

	var chats []PersonalChat = []PersonalChat{}

	err = cursor.All(ctx, &chats)
	if err != nil {
		return []PersonalChat{}, err
	}
	return chats, err
}
