package mongo

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
)

// MongoUseCase handles MongoDB operations business logic
type MongoUseCase struct {
	repo repository.MongoRepository
}

// NewMongoUseCase creates a new MongoDB use case
func NewMongoUseCase(repo repository.MongoRepository) *MongoUseCase {
	return &MongoUseCase{
		repo: repo,
	}
}

// GetDocuments retrieves documents with filter
func (uc *MongoUseCase) GetDocuments(ctx context.Context, query *entity.MongoQuery) ([]bson.M, error) {
	if query.Projection != nil {
		return uc.repo.FindWithProjection(ctx, query.Database, query.Collection, query.Filter, query.Projection, query.Limit)
	}
	return uc.repo.Find(ctx, query.Database, query.Collection, query.Filter, query.Limit)
}

// GetDocument retrieves a single document
func (uc *MongoUseCase) GetDocument(ctx context.Context, database, collection string, filter bson.M) (bson.M, error) {
	return uc.repo.FindOne(ctx, database, collection, filter)
}

// InsertDocument inserts a single document
func (uc *MongoUseCase) InsertDocument(ctx context.Context, insert *entity.MongoInsert) (interface{}, error) {
	return uc.repo.InsertOne(ctx, insert.Database, insert.Collection, insert.Document)
}

// InsertManyDocuments inserts multiple documents
func (uc *MongoUseCase) InsertManyDocuments(ctx context.Context, database, collection string, documents []interface{}) ([]interface{}, error) {
	return uc.repo.InsertMany(ctx, database, collection, documents)
}

// UpdateDocument updates a single document
func (uc *MongoUseCase) UpdateDocument(ctx context.Context, update *entity.MongoUpdate) error {
	return uc.repo.UpdateOne(ctx, update.Database, update.Collection, update.Filter, update.Update, update.Upsert)
}

// DeleteDocument deletes a single document
func (uc *MongoUseCase) DeleteDocument(ctx context.Context, del *entity.MongoDelete) error {
	return uc.repo.DeleteOne(ctx, del.Database, del.Collection, del.Filter)
}
