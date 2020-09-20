package pkg

import (
	"context"

	customLogHandle "golang-gin-sample/pkg/loghandle"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *MongoUpdateHelper) UpdateSingleDocumentFromMongo() {
	customLogHandle.LogInfo(
		"UpdateSingleDocumentFromMongo",
		"Update Single Document From Mongo",
		"Update Single Document From Mongo",
	)

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)
	filter := bson.M{
		"$and": c.FindData,
	}
	update := bson.M{"$set": bson.M{c.UpdateField: c.UpdateData}}
	opts := options.Update().SetUpsert(c.Upsert)

	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
		opts,
	)

	if err != nil {
		customLogHandle.ErrorHandle(
			"UpdateSingleDocumentFromMongo",
			"UpdateSingleDocumentFromMongo",
			"UpdateSingleDocumentFromMongo (Get Mongo Document error)",
			err,
		)
	}
}
