package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

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

func (c *MongoFind) FindDocument() ([]map[string]interface{}, error) {
	var (
		err    error
		result []map[string]interface{}
	)

	ctx := context.TODO()
	collection := c.Cl.Database(c.Db).Collection(c.Coll)
	filter := bson.M{
		"$and": c.FindData,
	}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Printf("Find Document Failed: %v\n", err)
	}
	err = cur.All(ctx, &result)

	return result, err
}

func (c *MongoInsert) InsertDocument() {
	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	_, err := collection.InsertOne(ctx, c.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Printf("Insert Document Failed: %v\n", err)
	}
}

func (c *MongoUpdate) UpdateDocument() {
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
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Printf("Update Document Failed: %v\n", err)
	}
}

func (c *MongoDelete) DeleteDocument() error {
	ctx := context.TODO()
	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	filter := bson.M{
		"$and": c.FindData,
	}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Printf("Delete Document Failed: %v\n", err)
	}

	return err
}
