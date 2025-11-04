package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// MongoRepository defines the interface for MongoDB operations
type MongoRepository interface {
	// Query operations
	Find(ctx context.Context, database, collection string, filter bson.M, limit int64) ([]bson.M, error)
	FindWithProjection(ctx context.Context, database, collection string, filter bson.M, projection interface{}, limit int64) ([]bson.M, error)
	FindOne(ctx context.Context, database, collection string, filter bson.M) (bson.M, error)

	// Insert operations
	InsertOne(ctx context.Context, database, collection string, document interface{}) (interface{}, error)
	InsertMany(ctx context.Context, database, collection string, documents []interface{}) ([]interface{}, error)

	// Update operations
	UpdateOne(ctx context.Context, database, collection string, filter bson.M, update interface{}, upsert bool) error
	UpdateMany(ctx context.Context, database, collection string, filter bson.M, update interface{}) error

	// Delete operations
	DeleteOne(ctx context.Context, database, collection string, filter bson.M) error
	DeleteMany(ctx context.Context, database, collection string, filter bson.M) error
}
