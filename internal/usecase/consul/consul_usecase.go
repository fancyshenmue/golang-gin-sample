package consul

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

// ConsulUseCase handles Consul operations business logic
type ConsulUseCase struct {
	repo repository.ConsulRepository
}

// NewConsulUseCase creates a new Consul use case
func NewConsulUseCase(repo repository.ConsulRepository) *ConsulUseCase {
	return &ConsulUseCase{
		repo: repo,
	}
}

// GetKey retrieves a value by key
func (uc *ConsulUseCase) GetKey(ctx context.Context, key string) (string, error) {
	return uc.repo.Get(ctx, key)
}

// ListKeys retrieves all keys with a prefix
func (uc *ConsulUseCase) ListKeys(ctx context.Context, prefix string) (map[string]string, error) {
	return uc.repo.List(ctx, prefix)
}

// PutKey stores a key-value pair
func (uc *ConsulUseCase) PutKey(ctx context.Context, key, value string) error {
	return uc.repo.Put(ctx, key, value)
}

// DeleteKey removes a key
func (uc *ConsulUseCase) DeleteKey(ctx context.Context, key string) error {
	return uc.repo.Delete(ctx, key)
}
