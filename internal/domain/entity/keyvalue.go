package entity

// KeyValue represents a key-value pair for distributed stores (etcd, Consul)
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// KeyValueQuery represents a query for key-value operations
type KeyValueQuery struct {
	Key    string `json:"key"`
	Prefix string `json:"prefix"`
}

// KeyValuePut represents a put operation
type KeyValuePut struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}
