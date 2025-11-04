package repository

import "context"

// EtcdRepository defines the interface for etcd operations
type EtcdRepository interface {
	Get(ctx context.Context, key string) (string, error)
	GetWithPrefix(ctx context.Context, prefix string) (map[string]string, error)
	Put(ctx context.Context, key, value string) error
	Delete(ctx context.Context, key string) error
}

// ConsulRepository defines the interface for Consul KV operations
type ConsulRepository interface {
	Get(ctx context.Context, key string) (string, error)
	List(ctx context.Context, prefix string) (map[string]string, error)
	Put(ctx context.Context, key, value string) error
	Delete(ctx context.Context, key string) error
}
