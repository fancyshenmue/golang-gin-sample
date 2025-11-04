package mongodb

import (
	"context"
	"fmt"

	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserMongoRepository implements repository.UserRepository using MongoDB
type UserMongoRepository struct {
	client   *mongo.Client
	database string
}

// NewUserMongoRepository creates a new user repository
func NewUserMongoRepository(client *mongo.Client, database string) *UserMongoRepository {
	return &UserMongoRepository{
		client:   client,
		database: database,
	}
}

// FindByUsername finds a user by username
func (r *UserMongoRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	collection := r.client.Database(r.database).Collection("users")

	var user entity.User
	filter := bson.M{"username": username}

	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("find user error: %w", err)
	}

	return &user, nil
}

// Create creates a new user
func (r *UserMongoRepository) Create(ctx context.Context, user *entity.User) error {
	collection := r.client.Database(r.database).Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}

	return nil
}

// Update updates a user
func (r *UserMongoRepository) Update(ctx context.Context, user *entity.User) error {
	collection := r.client.Database(r.database).Collection("users")

	filter := bson.M{"username": user.Username}
	update := bson.M{"$set": user}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update user error: %w", err)
	}

	return nil
}

// Delete deletes a user
func (r *UserMongoRepository) Delete(ctx context.Context, username string) error {
	collection := r.client.Database(r.database).Collection("users")

	filter := bson.M{"username": username}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}

	return nil
}
