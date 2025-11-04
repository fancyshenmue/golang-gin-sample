package mongodb

import (
	"context"
	"fmt"

	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GroupMongoRepository implements repository.GroupRepository using MongoDB
type GroupMongoRepository struct {
	client   *mongo.Client
	database string
}

// NewGroupMongoRepository creates a new group repository
func NewGroupMongoRepository(client *mongo.Client, database string) *GroupMongoRepository {
	return &GroupMongoRepository{
		client:   client,
		database: database,
	}
}

// FindByGroupName finds a group by name
func (r *GroupMongoRepository) FindByGroupName(ctx context.Context, groupName string) (*entity.Group, error) {
	collection := r.client.Database(r.database).Collection("groups")

	var group entity.Group
	filter := bson.M{"groupname": groupName}

	if err := collection.FindOne(ctx, filter).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("group not found")
		}
		return nil, fmt.Errorf("find group error: %w", err)
	}

	return &group, nil
}

// FindGroupsByMember finds all groups that contain a member
func (r *GroupMongoRepository) FindGroupsByMember(ctx context.Context, username string) ([]*entity.Group, error) {
	collection := r.client.Database(r.database).Collection("groups")

	filter := bson.M{"members": username}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find groups error: %w", err)
	}
	defer cur.Close(ctx)

	var groups []*entity.Group
	if err := cur.All(ctx, &groups); err != nil {
		return nil, fmt.Errorf("decode groups error: %w", err)
	}

	return groups, nil
}

// Create creates a new group
func (r *GroupMongoRepository) Create(ctx context.Context, group *entity.Group) error {
	collection := r.client.Database(r.database).Collection("groups")

	_, err := collection.InsertOne(ctx, group)
	if err != nil {
		return fmt.Errorf("create group error: %w", err)
	}

	return nil
}

// Update updates a group
func (r *GroupMongoRepository) Update(ctx context.Context, group *entity.Group) error {
	collection := r.client.Database(r.database).Collection("groups")

	filter := bson.M{"gid": group.GID}
	update := bson.M{"$set": group}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update group error: %w", err)
	}

	return nil
}

// Delete deletes a group
func (r *GroupMongoRepository) Delete(ctx context.Context, gid string) error {
	collection := r.client.Database(r.database).Collection("groups")

	filter := bson.M{"gid": gid}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete group error: %w", err)
	}

	return nil
}
