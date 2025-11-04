package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserNotInGroup     = errors.New("user not in required group")
)

// AuthUseCase handles authentication business logic
type AuthUseCase struct {
	userRepo  repository.UserRepository
	groupRepo repository.GroupRepository
}

// NewAuthUseCase creates a new authentication use case
func NewAuthUseCase(userRepo repository.UserRepository, groupRepo repository.GroupRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo:  userRepo,
		groupRepo: groupRepo,
	}
}

// Login authenticates a user with username and password
func (uc *AuthUseCase) Login(ctx context.Context, username, password string) (*entity.User, error) {
	// Find user by username
	user, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Validate password (MD5 hash)
	hashedPassword := hashPassword(password)
	if user.Password != hashedPassword {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// ValidateUserGroup checks if user belongs to a specific group
func (uc *AuthUseCase) ValidateUserGroup(ctx context.Context, username, requiredGroup string) error {
	groups, err := uc.groupRepo.FindGroupsByMember(ctx, username)
	if err != nil {
		return err
	}

	for _, group := range groups {
		if group.GroupName == requiredGroup {
			return nil
		}
	}

	return ErrUserNotInGroup
}

// hashPassword creates MD5 hash of password
func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
