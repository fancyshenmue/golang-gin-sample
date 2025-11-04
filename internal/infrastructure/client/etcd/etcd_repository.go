package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdClientRepository implements repository.EtcdRepository
type EtcdClientRepository struct {
	client *clientv3.Client
	kv     clientv3.KV
}

// NewEtcdClientRepository creates a new etcd repository
func NewEtcdClientRepository(endpoints []string, username, password string) (*EtcdClientRepository, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		Username:    username,
		Password:    password,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %w", err)
	}

	kv := clientv3.NewKV(cli)

	return &EtcdClientRepository{
		client: cli,
		kv:     kv,
	}, nil
}

// Get retrieves a value by key
func (r *EtcdClientRepository) Get(ctx context.Context, key string) (string, error) {
	resp, err := r.kv.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("key not found")
	}

	return string(resp.Kvs[0].Value), nil
}

// GetWithPrefix retrieves all keys with a prefix
func (r *EtcdClientRepository) GetWithPrefix(ctx context.Context, prefix string) (map[string]string, error) {
	resp, err := r.kv.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get keys with prefix: %w", err)
	}

	result := make(map[string]string)
	for _, kv := range resp.Kvs {
		result[string(kv.Key)] = string(kv.Value)
	}

	return result, nil
}

// Put stores a key-value pair
func (r *EtcdClientRepository) Put(ctx context.Context, key, value string) error {
	_, err := r.kv.Put(ctx, key, value)
	if err != nil {
		return fmt.Errorf("failed to put key: %w", err)
	}

	return nil
}

// Delete removes a key
func (r *EtcdClientRepository) Delete(ctx context.Context, key string) error {
	_, err := r.kv.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}

	return nil
}

// Close closes the etcd client connection
func (r *EtcdClientRepository) Close() error {
	return r.client.Close()
}
