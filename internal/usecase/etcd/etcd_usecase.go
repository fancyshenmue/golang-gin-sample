package etcd

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

// EtcdUseCase handles etcd operations business logic
type EtcdUseCase struct {
	repo repository.EtcdRepository
}

// NewEtcdUseCase creates a new etcd use case
func NewEtcdUseCase(repo repository.EtcdRepository) *EtcdUseCase {
	return &EtcdUseCase{
		repo: repo,
	}
}

// GetKey retrieves a value by key
func (uc *EtcdUseCase) GetKey(ctx context.Context, key string) (string, error) {
	return uc.repo.Get(ctx, key)
}

// GetKeysWithPrefix retrieves all keys with a prefix
func (uc *EtcdUseCase) GetKeysWithPrefix(ctx context.Context, prefix string) (map[string]string, error) {
	return uc.repo.GetWithPrefix(ctx, prefix)
}

// PutKey stores a key-value pair
func (uc *EtcdUseCase) PutKey(ctx context.Context, key, value string) error {
	return uc.repo.Put(ctx, key, value)
}

// DeleteKey removes a key
func (uc *EtcdUseCase) DeleteKey(ctx context.Context, key string) error {
	return uc.repo.Delete(ctx, key)
}
