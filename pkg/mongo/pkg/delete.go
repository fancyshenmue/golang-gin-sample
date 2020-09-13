package pkg

import (
	"context"

	customLogHandle "golang-gin-sample/pkg/loghandle"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteOneDataFromMongo | Delete One Data From Mongo
func (c *MongoResultHelper) DeleteOneDataFromMongo() error {
	customLogHandle.LogInfo(
		"DeleteOneDataFromMongo",
		"DeleteOneDataFromMongo",
		"DeleteOneDataFromMongo",
	)

	ctx := context.TODO()
	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{c.FindField: c.FindData}
	_, err := collection.DeleteOne(ctx, filter)
	customLogHandle.ErrorHandle(
		"DeleteOneDataFromMongo",
		"DeleteOneDataFromMongo",
		"DeleteOneDataFromMongo error",
		err,
	)

	return err
}

// DeleteDataComplexFromMongo | Find Data Complex Query From Mongo
func (c *MongoResultComplexQueryHelper) DeleteDataComplexFromMongo() error {
	customLogHandle.LogInfo(
		"DeleteDataComplexFromMongo",
		"Delete Data Complex From Mongo",
		"Delete Data Complex From Mongo",
	)

	ctx := context.TODO()
	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{
		"$and": c.FindData,
	}

	_, err := collection.DeleteOne(ctx, filter)
	customLogHandle.ErrorHandle(
		"DeleteDataComplexFromMongo",
		"Delete Data Complex From Mongo",
		"Delete Data Complex From Mongo filter error",
		err,
	)

	return err
}
