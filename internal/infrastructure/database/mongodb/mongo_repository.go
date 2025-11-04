package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBRepository implements repository.MongoRepository
type MongoDBRepository struct {
	client *mongo.Client
}

// NewMongoDBRepository creates a new MongoDB repository
func NewMongoDBRepository(client *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{
		client: client,
	}
}

// Find retrieves documents matching the filter
func (r *MongoDBRepository) Find(ctx context.Context, database, collection string, filter bson.M, limit int64) ([]bson.M, error) {
	coll := r.client.Database(database).Collection(collection)

	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(limit)
	}

	cur, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("find error: %w", err)
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return results, nil
}

// FindWithProjection retrieves documents with projection
func (r *MongoDBRepository) FindWithProjection(ctx context.Context, database, collection string, filter bson.M, projection interface{}, limit int64) ([]bson.M, error) {
	coll := r.client.Database(database).Collection(collection)

	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(limit)
	}
	if projection != nil {
		opts.SetProjection(projection)
	}

	cur, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("find error: %w", err)
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return results, nil
}

// FindOne retrieves a single document
func (r *MongoDBRepository) FindOne(ctx context.Context, database, collection string, filter bson.M) (bson.M, error) {
	coll := r.client.Database(database).Collection(collection)

	var result bson.M
	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("document not found")
		}
		return nil, fmt.Errorf("find one error: %w", err)
	}

	return result, nil
}

// InsertOne inserts a single document
func (r *MongoDBRepository) InsertOne(ctx context.Context, database, collection string, document interface{}) (interface{}, error) {
	coll := r.client.Database(database).Collection(collection)

	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("insert one error: %w", err)
	}

	return result.InsertedID, nil
}

// InsertMany inserts multiple documents
func (r *MongoDBRepository) InsertMany(ctx context.Context, database, collection string, documents []interface{}) ([]interface{}, error) {
	coll := r.client.Database(database).Collection(collection)

	result, err := coll.InsertMany(ctx, documents)
	if err != nil {
		return nil, fmt.Errorf("insert many error: %w", err)
	}

	return result.InsertedIDs, nil
}

// UpdateOne updates a single document
func (r *MongoDBRepository) UpdateOne(ctx context.Context, database, collection string, filter bson.M, update interface{}, upsert bool) error {
	coll := r.client.Database(database).Collection(collection)

	opts := options.Update().SetUpsert(upsert)

	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("update one error: %w", err)
	}

	return nil
}

// UpdateMany updates multiple documents
func (r *MongoDBRepository) UpdateMany(ctx context.Context, database, collection string, filter bson.M, update interface{}) error {
	coll := r.client.Database(database).Collection(collection)

	_, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update many error: %w", err)
	}

	return nil
}

// DeleteOne deletes a single document
func (r *MongoDBRepository) DeleteOne(ctx context.Context, database, collection string, filter bson.M) error {
	coll := r.client.Database(database).Collection(collection)

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one error: %w", err)
	}

	return nil
}

// DeleteMany deletes multiple documents
func (r *MongoDBRepository) DeleteMany(ctx context.Context, database, collection string, filter bson.M) error {
	coll := r.client.Database(database).Collection(collection)

	_, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete many error: %w", err)
	}

	return nil
}

// Connect creates a new MongoDB connection
func Connect(mongoURI string, credential *options.Credential) (*mongo.Client, error) {
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(mongoURI)
	if credential != nil {
		clientOptions.SetAuth(*credential)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connect error: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return client, nil
}
