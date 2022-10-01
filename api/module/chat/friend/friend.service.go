package friend

import (
	"context"
	"web-chat/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBMongo              = "chat"
	CollectionFriendship = "friend"
)

type FriendshipService interface {
	createFriendship(context.Context, Friendship) error
	getFriendship(context.Context, string) ([]Friendship, error)
	deleteFriendship(context.Context, string) error
}

type friendshipService struct {
	MongoDB *mongo.Client
}

func NewFriendshipService(datasource database.DataSource) FriendshipService {
	return &friendshipService{
		MongoDB: datasource.MongoDB,
	}
}

func (fs *friendshipService) createFriendship(ctx context.Context, friendship Friendship) error {
	friendshipCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionFriendship)

	return friendshipCollection.Database().Client().UseSession(ctx, func(sc mongo.SessionContext) error {
		err := sc.StartTransaction()
		if err != nil {
			return err
		}

		_, err = friendshipCollection.InsertOne(sc, friendship)
		if err != nil {
			return err
		}

		friendship.MeId, friendship.FriendId = friendship.FriendshipId, friendship.MeId
		_, err = friendshipCollection.InsertOne(sc, friendship)
		if err != nil {
			return err
		}

		return nil
	})

}

func (fs *friendshipService) getFriendship(ctx context.Context, meId string) ([]Friendship, error) {

	friendshipCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionFriendship)

	find := bson.M{
		"meId":      meId,
		"isDeleted": false,
	}

	cursor, err := friendshipCollection.Find(ctx, find)
	if err != nil {
		return []Friendship{}, err
	}

	var friends []Friendship

	err = cursor.Decode(&friends)
	if err != nil {
		return []Friendship{}, err
	}

	return friends, nil

}

func (fs *friendshipService) deleteFriendship(ctx context.Context, friendshipId string) error {
	friendshipCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionFriendship)

	find := bson.M{
		"friendshipId": friendshipId,
	}

	update := bson.M{
		"$set": bson.M{
			"isDeleted": false,
		},
	}

	err := friendshipCollection.FindOneAndUpdate(ctx, find, update).Err()
	if err != nil {
		return err
	}
	return nil
}
