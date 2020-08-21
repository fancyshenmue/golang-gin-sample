package pkg

import (
	"context"
	"fmt"

	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneDataFromMongo | Find One Data From Mongo
func (c *MongoResultHelper) FindOneDataFromMongo() ([]bson.M, error) {
	customLogHandle.LogInfo(
		"FindOneDataFromMongo",
		"select one data from mongo",
		"select one data from mongo",
	)

	var result []bson.M

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{c.FindField: c.FindData}

	cur, err := collection.Find(ctx, filter)
	customLogHandle.ErrorHandle(
		"FindOneDataFromMongo",
		"select one data from mongo",
		"select one data from mongo filter error",
		err,
	)

	err = cur.All(ctx, &result)
	customLogHandle.ErrorHandle(
		"FindOneDataFromMongo",
		"select one data from mongo",
		"select one data from mongo result error",
		err,
	)

	return result, err
}

// FindDataComplexQueryFromMongo | Find Data Complex Query From Mongo
func (c *MongoResultComplexQueryHelper) FindDataComplexQueryFromMongo() ([]bson.M, error) {
	customLogHandle.LogInfo(
		"FindDataComplexQueryFromMongo",
		"Find Data Complex Query From Mongo",
		"Find Data Complex Query From Mongo",
	)

	var (
		err    error
		result []bson.M
	)

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{
		"$and": c.FindData,
	}

	cur, err := collection.Find(ctx, filter)
	customLogHandle.ErrorHandle(
		"FindDataComplexQueryFromMongo",
		"Find Data Complex Query From Mongo",
		"Find Data Complex Query From Mongo filter error",
		err,
	)

	err = cur.All(ctx, &result)
	customLogHandle.ErrorHandle(
		"FindDataComplexQueryFromMongo",
		"Find Data Complex Query From Mongo",
		"Find Data Complex Query From Mongo result error",
		err,
	)

	return result, err
}

// FindDataRegexFromMongo | Find Data Regex From Mongo
func (c *MongoResultHelper) FindDataRegexFromMongo() []bson.M {
	customLogHandle.LogInfo(
		"FindDataRegexFromMongo",
		"select data (regex) from mongo",
		"select data (regex) from mongo",
	)

	var episodesFiltered []bson.M

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.D{{"name", primitive.Regex{Pattern: c.FindData, Options: ""}}}

	cur, err := collection.Find(ctx, filter)
	customLogHandle.ErrorHandle(
		"FindDataRegexFromMongo",
		"select data (regex) from mongo",
		"select data (regex) from mongo filter error",
		err,
	)

	cur.All(ctx, &episodesFiltered)
	customLogHandle.ErrorHandle(
		"FindDataRegexFromMongo",
		"select data (regex) from mongo",
		"select data (regex) from mongo result error",
		err,
	)

	return episodesFiltered
}

// ResponseSample | Response Sample
func ResponseSample(client *mongo.Client, databaseName, collectionName string, findNum int64) []bson.M {
	customLogHandle.LogInfo(
		"ResponseSample",
		"select all white label info from mongo",
		"select all white label info from mongo",
	)

	var result []bson.M

	ctx := context.TODO()

	collection := client.Database(databaseName).Collection(collectionName)

	findOptions := options.Find()
	findOptions.SetLimit(findNum)

	cur, err := collection.Find(ctx, bson.M{}, findOptions)
	defer cur.Close(ctx)
	customLogHandle.ErrorHandle(
		"ResponseSample",
		"select all white label info from mongo",
		"select all white label info from mongo cursor error",
		err,
	)

	for cur.Next(ctx) {
		var elem bson.M
		err := cur.Decode(&elem)
		customLogHandle.ErrorHandle(
			"ResponseSample",
			"select all white label info from mongo",
			"select all white label info from mongo cursor append error",
			err,
		)

		result = append(result, elem)
	}

	err = cur.Err()
	customLogHandle.ErrorHandle(
		"ResponseSample",
		"select all white label info from mongo",
		"select all white label info from mongo result error",
		err,
	)

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)

	return result
}

// FindGroupInfoFromMongo | Find Group Info From Mongo
func (c *MongoResultHelper) FindGroupInfoFromMongo() ([]AuthMemberTop, error) {
	customLogHandle.LogInfo(
		"FindGroupInfoFromMongo",
		"Find Group Info From Mongo",
		"Find Group Info From Mongo",
	)

	var result []AuthMemberTop

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{c.FindField: c.FindData}

	cur, err := collection.Find(ctx, filter)
	customLogHandle.ErrorHandle(
		"FindGroupInfoFromMongo",
		"Find Group Info From Mongo",
		"Find Group Info From Mongo filter error",
		err,
	)

	err = cur.All(ctx, &result)
	customLogHandle.ErrorHandle(
		"FindGroupInfoFromMongo",
		"Find Group Info From Mongo",
		"Find Group Info From Mongo result error",
		err,
	)

	return result, err
}
