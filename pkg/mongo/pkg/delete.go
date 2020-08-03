package pkg

import (
	"context"

	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteOneDataFromMongo | Delete One Data From Mongo
func (c *MongoResultHelper) DeleteOneDataFromMongo() error {
	customLogHandle.LogInfo(
		"DeleteOneDataFromMongo",
		"white label info delete",
		"white label info delete",
	)

	ctx := context.TODO()
	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{c.FindField: c.FindData}
	_, err := collection.DeleteOne(ctx, filter)
	customLogHandle.ErrorHandle(
		"WhiteLabelInfoInsertOne",
		"white label info delete",
		"white label info delete error",
		err,
	)

	return err
}
