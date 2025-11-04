# Golang Gin Sample - Clean Architecture

A unified RESTful API built with Go, Gin framework, and Clean Architecture principles.

## Overview

This project demonstrates a production-ready microservices architecture following Clean Architecture principles. It provides a unified API with support for MongoDB, Etcd, Consul, and CDN operations.

## Architecture

This application follows **Clean Architecture** with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                    DOMAIN LAYER                         │
│          (Entities, Repository Interfaces)              │
│                  NO DEPENDENCIES                        │
└─────────────────────┬───────────────────────────────────┘
                      ↑
┌─────────────────────┴───────────────────────────────────┐
│                   USE CASE LAYER                        │
│           (Application Business Logic)                  │
└─────────────────────┬───────────────────────────────────┘
                      ↑
┌─────────────────────┴───────────────────────────────────┐
│                   ADAPTER LAYER                         │
│         (HTTP Handlers, Presenters)                     │
└─────────────────────┬───────────────────────────────────┘
                      ↑
┌─────────────────────┴───────────────────────────────────┐
│                INFRASTRUCTURE LAYER                     │
│     (Database, External APIs, Frameworks)               │
└─────────────────────────────────────────────────────────┘
```

See [CLEAN_ARCHITECTURE.md](./CLEAN_ARCHITECTURE.md) for detailed documentation.

## Features

- ✅ **Clean Architecture** - Maintainable, testable, and scalable
- ✅ **Unified API** - Single service with multiple endpoints
- ✅ **MongoDB Support** - Document database operations
- ✅ **Etcd Support** - Distributed configuration management
- ✅ **Consul Support** - Service discovery and KV store
- ✅ **JWT Authentication** - Secure token-based auth
- ✅ **Swagger Documentation** - Interactive API docs
- ✅ **Docker Ready** - Containerized deployment
- ✅ **Kubernetes Ready** - K8s deployment with Kustomize

## Directory Structure

```
.
├── cmd/api/                    # Application entry point
│   └── main.go                 # Unified API server
├── internal/                   # Private application code
│   ├── domain/                 # Domain layer (entities & interfaces)
│   ├── usecase/                # Use case layer (business logic)
│   ├── adapter/                # Adapter layer (handlers & middleware)
│   └── infrastructure/         # Infrastructure layer (DB, clients, config)
├── pkg/                        # Public shared utilities
│   ├── errorhandle/           # Error handling utilities
│   ├── http/                  # HTTP utilities
│   ├── loghandle/             # Logging utilities
│   └── validation/            # Input validation
├── api/                        # OpenAPI/Swagger specs
├── build/                      # Build configurations
├── deployments/                # Deployment manifests
└── migrations/                 # Database migrations
```

## Quick Start

### Prerequisites

- Go 1.16+
- MongoDB (optional - for MongoDB endpoints)
- Etcd (optional - for Etcd endpoints)
- Consul (optional - for Consul endpoints)

### Environment Variables

```bash
# MongoDB Configuration (required for /api/v1/mongo endpoints)
export MONGO_ADDRESS="localhost:27017"
export MONGO_USER="admin"
export MONGO_PASS="password"
export MONGO_DATABASE="test"

# JWT Configuration (required)
export JWT_KEY="your-secret-key"
export JWT_REALM="api"
export JWT_TOKEN_HEADER="Bearer"

# Etcd Configuration (optional)
export ETCD_ENDPOINTS="localhost:2379"
export ETCD_CONNECT_USER=""
export ETCD_CONNECT_PASSWD=""

# Consul Configuration (optional)
export CONSUL_ADDRESS="localhost:8500"
export CONSUL_TOKEN=""

# Server Configuration
export APP_PORT="8080"
```

### Run Locally

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/api/main.go
```

The server will start on http://localhost:8080

### API Documentation

Visit http://localhost:8080/swagger/index.html for interactive API documentation.

### Health Check

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "healthy",
  "services": {
    "mongodb": "connected",
    "etcd": "available",
    "consul": "available"
  }
}
```

## API Endpoints

### Authentication

All API endpoints (except `/auth/*`, `/health`, and `/`) require JWT authentication.

```bash
# Login
POST /auth/login
{
  "username": "user",
  "password": "pass"
}

# Refresh Token
GET /auth/refresh_token

# Logout
POST /auth/logout
```

### MongoDB Operations

```bash
# Get Documents
POST /api/v1/mongo/getDocument
{
  "database": "test",
  "collection": "users",
  "filter": {"status": "active"},
  "limit": 10
}

# Insert Single Document
POST /api/v1/mongo/insertSingleDocument

# Insert Many Documents
POST /api/v1/mongo/insertManyDocument

# Update Document
PUT /api/v1/mongo/updateSingleDocument

# Delete Document
DELETE /api/v1/mongo/deleteSingleDocument
```

### Etcd Operations

```bash
# Get Key
POST /api/v1/etcd/getKey

# Get Keys with Prefix
POST /api/v1/etcd/getKeyWithPrefix

# Put Key
POST /api/v1/etcd/putKey

# Delete Key
DELETE /api/v1/etcd/deleteKey
```

### Consul Operations

```bash
# Get Key
POST /api/v1/consul/getKey

# List Keys
POST /api/v1/consul/listKey

# Put Key
POST /api/v1/consul/putKey

# Delete Key
DELETE /api/v1/consul/deleteKey
```

## Docker

### Build Docker Image

```bash
# Build binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api cmd/api/main.go

# Build image
docker build -t golang-gin-sample:latest -f build/Docker/api/Dockerfile .
```

### Run with Docker

```bash
docker run -p 8080:8080 \
  -e MONGO_ADDRESS="mongodb:27017" \
  -e JWT_KEY="secret" \
  golang-gin-sample:latest
```

### Docker Compose

```bash
cd deployments/docker-compose
docker-compose up -d
```

## Kubernetes Deployment

```bash
# Deploy with Kustomize
kubectl apply -k deployments/Kustomize/api/

# Check status
kubectl get pods
kubectl get svc
```

## Development

### Project Structure

- **Domain Layer** (`internal/domain/`) - Core business logic, entities, and repository interfaces
- **Use Case Layer** (`internal/usecase/`) - Application-specific business rules
- **Adapter Layer** (`internal/adapter/`) - HTTP handlers, middleware, presenters
- **Infrastructure Layer** (`internal/infrastructure/`) - Database implementations, external clients

### Adding a New Feature

1. **Define Entity** in `internal/domain/entity/`
2. **Define Repository Interface** in `internal/domain/repository/`
3. **Implement Use Case** in `internal/usecase/`
4. **Implement Repository** in `internal/infrastructure/`
5. **Create Handler** in `internal/adapter/handler/http/`
6. **Register Routes** in `cmd/api/main.go`

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/usecase/auth/...
```

## Benefits of Clean Architecture

✅ **Testability** - Business logic can be tested without HTTP or database
✅ **Flexibility** - Easy to swap database or framework implementations
✅ **Maintainability** - Clear separation of concerns
✅ **Scalability** - Well-organized structure for growing codebase
✅ **Independence** - Domain logic is independent of frameworks

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Author

Charles Hsu (fancyshenmue@gmail.com)

## Documentation

- [Clean Architecture Guide](./CLEAN_ARCHITECTURE.md) - Comprehensive architecture documentation
- [API Documentation](http://localhost:8080/swagger/index.html) - Swagger UI (when server is running)

## Support

For issues, questions, or contributions, please open an issue on GitHub.
