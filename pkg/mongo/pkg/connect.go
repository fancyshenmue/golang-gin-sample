package pkg

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongo | Connect To Mongo
func ConnectToMongo(mongoURI string, credential options.Credential) (*mongo.Client, error) {
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("ping error: %v\n", err)
	}

	return client, err
}
