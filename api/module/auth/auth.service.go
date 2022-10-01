package auth

import (
	"context"
	"fmt"
	"web-chat/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBMongo        = "chat"
	CollectionAuth = "auth"
)

type AuthService interface {
	checkAuth(context.Context, string) (Auth, error)
	register(context.Context, Auth) error
	login(context.Context, Auth) error
	changePassword(context.Context, Auth) error
}

type authService struct {
	MongoDB *mongo.Client
}

func NewAuthService(datasource database.DataSource) AuthService {
	return &authService{
		MongoDB: datasource.MongoDB,
	}
}

func (as *authService) checkAuth(ctx context.Context, username string) (Auth, error) {
	authCollection := as.MongoDB.Database(DBMongo).Collection(CollectionAuth)

	query := bson.M{
		"username": username,
	}

	var auth Auth
	err := authCollection.FindOne(ctx, query).Decode(&auth)
	if err != nil {
		return auth, fmt.Errorf("account not found")
	}
	return auth, nil
}

func (as *authService) register(ctx context.Context, auth Auth) error {

	authCollection := as.MongoDB.Database(DBMongo).Collection(CollectionAuth)

	_, err := authCollection.InsertOne(ctx, auth)

	if err != nil {
		return err
	}

	return nil

}
func (as *authService) login(ctx context.Context, auth Auth) error {

	authCollection := as.MongoDB.Database(DBMongo).Collection(CollectionAuth)

	query := bson.M{
		"username": auth.Username,
	}

	var dbAuth Auth
	err := authCollection.FindOne(ctx, query).Decode(&dbAuth)
	if err != nil {
		return fmt.Errorf("account not found")
	}

	if err = comparePassword(dbAuth.Password, auth.Password); err != nil {
		return fmt.Errorf("invalid credentails")
	}

	return nil

}
func (as *authService) changePassword(ctx context.Context, auth Auth) error {

	authCollection := as.MongoDB.Database(DBMongo).Collection(CollectionAuth)

	find := bson.M{
		"username": auth.Username,
	}

	update := bson.M{
		"$set": bson.M{
			"password": auth.Password,
		},
	}

	err := authCollection.FindOneAndUpdate(ctx, find, update).Decode(&auth)
	if err != nil {
		return err
	}

	return nil
}
