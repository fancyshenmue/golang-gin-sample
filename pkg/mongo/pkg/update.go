package pkg

import (
	"context"

	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *MongoUpdateHelper) UpdateSingleDocumentFromMongo() {
	customLogHandle.LogInfo(
		"UpdateSingleDocumentFromMongo",
		"Update Single Document From Mongo",
		"Update Single Document From Mongo",
	)

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)
	filter := bson.M{"name": bson.M{"$eq": c.FindData}}
	update := bson.M{"$set": bson.M{c.Field: c.UpdateData}}
	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
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
