package bootstrap

import (
	"context"
	"log"
	"lynx/internal/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Db.Mongo.Uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	logger.Info("Mongodb has been initialize")
	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	logger.Info("Connection to MongoDB closed.")
}
