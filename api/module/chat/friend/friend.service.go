package friend

import (
	"context"
	"fmt"
	"web-chat/database"
	"web-chat/pkg/mongoquerybuilder"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBMongo              = "chat"
	CollectionFriendship = "friend"
	CollectionAuth       = "auth"
)

type FriendshipService interface {
	createFriendship(context.Context, Friendship) error
	getFriendship(context.Context, string) ([]FriendshipResult, error)
	deleteFriendship(context.Context, DeleteFriendshipDTO) error
	searchFriendship(context.Context, SearchFriendshipDTO, string) ([]Search, error)
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

	query := mongoquerybuilder.Query{
		Query: `{
			"$and": [{ "meId": "~1" }, { "friendId": "~2" }],
			"isDeleted": false
		  }`,
		Params: mongoquerybuilder.Params{friendship.MeId, friendship.FriendId},
	}

	find, err := query.QueryBuilder()
	if err != nil {
		return fmt.Errorf("invalid loader qeury")
	}

	cursor, err := friendshipCollection.Find(ctx, find.Query)
	if err != nil {
		return err
	}

	var dupliateEntry []Friendship = []Friendship{}
	cursor.All(ctx, &dupliateEntry)

	if len(dupliateEntry) > 0 {
		return fmt.Errorf("already friend")
	}

	_, err = friendshipCollection.InsertOne(ctx, friendship)
	if err != nil {
		return err
	}

	friendship.MeId, friendship.FriendId = friendship.FriendId, friendship.MeId
	_, err = friendshipCollection.InsertOne(ctx, friendship)
	if err != nil {
		return err
	}

	return nil
}

func (fs *friendshipService) getFriendship(ctx context.Context, meId string) ([]FriendshipResult, error) {

	friendshipCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionFriendship)

	query := mongoquerybuilder.Query{
		Query: `[
			{
			  "$match": {
				"meId": "~1",
				"isDeleted": false
			  }
			},
			{
			  "$lookup": {
				"from": "auth",
				"localField": "friendId",
				"foreignField": "authId",
				"as": "info"
			  }
			},
			{
			  "$unwind": "$info"
			},
			{
			  "$addFields": {
				"friendName": "$info.username"
			  }
			}
		  ]`,
		Params: mongoquerybuilder.Params{meId},
	}

	find, _ := query.QueryBuilder()

	cursor, err := friendshipCollection.Aggregate(ctx, find.Query)
	if err != nil {
		return []FriendshipResult{}, err
	}

	var friends []FriendshipResult = []FriendshipResult{}

	err = cursor.All(ctx, &friends)
	if err != nil {
		return []FriendshipResult{}, err
	}

	return friends, nil

}

func (fs *friendshipService) deleteFriendship(ctx context.Context, deleteFriendshipDTO DeleteFriendshipDTO) error {
	friendshipCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionFriendship)

	find := bson.M{
		"friendshipId": deleteFriendshipDTO.FriendshipId,
	}

	update := bson.M{
		"$set": bson.M{
			"isDeleted": true,
		},
	}

	_, err := friendshipCollection.UpdateMany(ctx, find, update)
	if err != nil {
		return err
	}

	return nil
}

func (fs *friendshipService) searchFriendship(ctx context.Context, searchFriendshipDTO SearchFriendshipDTO, myusername string) ([]Search, error) {
	authCollection := fs.MongoDB.Database(DBMongo).Collection(CollectionAuth)

	find := bson.M{
		"username": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + searchFriendshipDTO.Filter + ".*",
				Options: "i",
			},
		},
	}

	cursor, err := authCollection.Find(ctx, find)
	if err != nil {
		return []Search{}, err
	}

	var searchlist []Search = []Search{}

	err = cursor.All(ctx, &searchlist)
	if err != nil {
		return []Search{}, err
	}

	return searchlist, nil
}
