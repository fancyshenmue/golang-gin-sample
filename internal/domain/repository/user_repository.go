package repository

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, username string) error
}

// GroupRepository defines the interface for group data operations
type GroupRepository interface {
	FindByGroupName(ctx context.Context, groupName string) (*entity.Group, error)
	FindGroupsByMember(ctx context.Context, username string) ([]*entity.Group, error)
	Create(ctx context.Context, group *entity.Group) error
	Update(ctx context.Context, group *entity.Group) error
	Delete(ctx context.Context, gid string) error
}
