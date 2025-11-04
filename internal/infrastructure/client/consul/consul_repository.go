package consul

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
)

// ConsulClientRepository implements repository.ConsulRepository
type ConsulClientRepository struct {
	client *api.Client
	kv     *api.KV
}

// NewConsulClientRepository creates a new Consul repository
func NewConsulClientRepository(address, token string) (*ConsulClientRepository, error) {
	config := &api.Config{
		Address: address,
		Token:   token,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to consul: %w", err)
	}

	return &ConsulClientRepository{
		client: client,
		kv:     client.KV(),
	}, nil
}

// Get retrieves a value by key
func (r *ConsulClientRepository) Get(ctx context.Context, key string) (string, error) {
	pair, _, err := r.kv.Get(key, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}

	if pair == nil {
		return "", fmt.Errorf("key not found")
	}

	return string(pair.Value), nil
}

// List retrieves all keys with a prefix
func (r *ConsulClientRepository) List(ctx context.Context, prefix string) (map[string]string, error) {
	pairs, _, err := r.kv.List(prefix, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %w", err)
	}

	result := make(map[string]string)
	for _, pair := range pairs {
		result[pair.Key] = string(pair.Value)
	}

	return result, nil
}

// Put stores a key-value pair
func (r *ConsulClientRepository) Put(ctx context.Context, key, value string) error {
	p := &api.KVPair{
		Key:   key,
		Value: []byte(value),
	}

	_, err := r.kv.Put(p, nil)
	if err != nil {
		return fmt.Errorf("failed to put key: %w", err)
	}

	return nil
}

// Delete removes a key
func (r *ConsulClientRepository) Delete(ctx context.Context, key string) error {
	_, err := r.kv.Delete(key, nil)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}

	return nil
}
