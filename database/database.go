package database

import (
	"context"
	"fmt"
	"sync"
	"time"
	"web-chat/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

type DataSource struct {
	MongoDB *mongo.Client
}

func InitDatabase() (*DataSource, error) {

	mongoDNSUrl := fmt.Sprintf(config.DatabaseConfig.Mongo.DSN)
	mongodb, err := initMongoDB(mongoDNSUrl)
	if err != nil {
		return nil, err
	}

	return &DataSource{MongoDB: mongodb}, nil
}

func initMongoDB(mongoDNS string) (*mongo.Client, error) {

	var instance *mongo.Client
	var mongoerror error

	once.Do(func() {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoDNS)
		clientOptions.SetMaxConnIdleTime(100)
		clientOptions.SetMaxPoolSize(1000)
		clientOptions.SetMaxConnIdleTime(4 * time.Hour)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			instance = nil
			mongoerror = err
		}

		if err = client.Ping(ctx, nil); err != nil {
			instance = nil
			mongoerror = err
		}

		instance = client
		mongoerror = nil
	})

	return instance, mongoerror

}
