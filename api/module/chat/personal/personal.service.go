package personal

import (
	"context"
	"web-chat/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBMongo        = "chat"
	CollectionChat = "chat"
)

type PersonalChatService interface {
	createChat(context.Context, PersonalChat) error
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
