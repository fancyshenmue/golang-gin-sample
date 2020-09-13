package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	customLogHandle "golang-gin-sample/pkg/loghandle"
)

// InsertSingleDocument | Insert Single Document
func InsertSingleDocument(client *mongo.Client, databaseName, collectionName string, data bson.M) {
	customLogHandle.LogInfo(
		"InsertSingleDocument",
		"Insert Single Document",
		"Insert Single Document",
	)

	ctx := context.TODO()

	collection := client.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	customLogHandle.ErrorHandle(
		"InsertSingleDocument",
		"Insert Single Document",
		"Insert Single Document error",
		err,
	)
}
