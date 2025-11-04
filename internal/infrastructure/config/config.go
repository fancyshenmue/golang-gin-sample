package config

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	MongoDB  MongoDBConfig
	JWT      JWTConfig
	Etcd     EtcdConfig
	Consul   ConsulConfig
	CDN      CDNConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// MongoDBConfig holds MongoDB configuration
type MongoDBConfig struct {
	Address    string
	User       string
	Password   string
	Database   string
	URI        string
	Credential options.Credential
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey   string
	Realm       string
	TokenHeader string
	Timeout     int
}

// EtcdConfig holds etcd configuration
type EtcdConfig struct {
	Endpoints []string
}

// ConsulConfig holds Consul configuration
type ConsulConfig struct {
	Address string
	Scheme  string
}

// CDNConfig holds CDN provider configurations
type CDNConfig struct {
	Cloudflare CloudflareConfig
	Verizon    VerizonConfig
	GreyCDN    GreyCDNConfig
}

// CloudflareConfig holds Cloudflare configuration
type CloudflareConfig struct {
	APIKey    string
	APIEmail  string
	APIToken  string
}

// VerizonConfig holds Verizon configuration
type VerizonConfig struct {
	AccountNumber string
	APIToken      string
}

// GreyCDNConfig holds GreyCDN configuration
type GreyCDNConfig struct {
	APIKey string
	APISecret string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	mongoAddress := getEnv("MONGO_ADDRESS", "localhost:27017")
	mongoUser := getEnv("MONGO_USER", "")
	mongoPass := getEnv("MONGO_PASS", "")
	mongoDatabase := getEnv("MONGO_DATABASE", "test")

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		MongoDB: MongoDBConfig{
			Address:  mongoAddress,
			User:     mongoUser,
			Password: mongoPass,
			Database: mongoDatabase,
			URI:      fmt.Sprintf("mongodb://%s", mongoAddress),
			Credential: options.Credential{
				Username: mongoUser,
				Password: mongoPass,
			},
		},
		JWT: JWTConfig{
			SecretKey:   getEnv("JWT_KEY", "secret"),
			Realm:       getEnv("JWT_REALM", "test zone"),
			TokenHeader: getEnv("JWT_TOKEN_HEADER", "Bearer"),
			Timeout:     3600, // 1 hour
		},
		Etcd: EtcdConfig{
			Endpoints: []string{getEnv("ETCD_ENDPOINTS", "localhost:2379")},
		},
		Consul: ConsulConfig{
			Address: getEnv("CONSUL_ADDRESS", "localhost:8500"),
			Scheme:  getEnv("CONSUL_SCHEME", "http"),
		},
		CDN: CDNConfig{
			Cloudflare: CloudflareConfig{
				APIKey:   getEnv("CLOUDFLARE_API_KEY", ""),
				APIEmail: getEnv("CLOUDFLARE_API_EMAIL", ""),
				APIToken: getEnv("CLOUDFLARE_API_TOKEN", ""),
			},
			Verizon: VerizonConfig{
				AccountNumber: getEnv("VERIZON_ACCOUNT_NUMBER", ""),
				APIToken:      getEnv("VERIZON_API_TOKEN", ""),
			},
			GreyCDN: GreyCDNConfig{
				APIKey:    getEnv("GREYCDN_API_KEY", ""),
				APISecret: getEnv("GREYCDN_API_SECRET", ""),
			},
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
