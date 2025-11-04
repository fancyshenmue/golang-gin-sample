# Clean Architecture Refactoring

This document describes the Clean Architecture refactoring of the golang-gin-sample project.

## Table of Contents

- [Overview](#overview)
- [Architecture Layers](#architecture-layers)
- [Directory Structure](#directory-structure)
- [Dependency Flow](#dependency-flow)
- [Usage Guide](#usage-guide)
- [Migration Guide](#migration-guide)

## Overview

This project has been refactored to follow Clean Architecture principles, which provides:

- **Separation of Concerns**: Each layer has a specific responsibility
- **Testability**: Business logic can be tested independently
- **Maintainability**: Changes in one layer don't affect others
- **Flexibility**: Easy to swap implementations (e.g., change database)
- **Scalability**: Clear structure for growing codebase

## Architecture Layers

Clean Architecture organizes code into four concentric layers:

### 1. Domain Layer (Core)

**Location**: `internal/domain/`

**Purpose**: Contains core business entities and business rules. This is the heart of the application and has NO dependencies on external frameworks or libraries.

**Components**:
- `entity/` - Domain entities (User, Group, Document, etc.)
- `repository/` - Repository interfaces (ports)

**Example**:
```go
// internal/domain/entity/user.go
type User struct {
    Username  string
    Password  string
    FirstName string
    LastName  string
}

// internal/domain/repository/user_repository.go
type UserRepository interface {
    FindByUsername(ctx context.Context, username string) (*entity.User, error)
    Create(ctx context.Context, user *entity.User) error
}
```

### 2. Use Case Layer (Application)

**Location**: `internal/usecase/`

**Purpose**: Contains application-specific business rules. Orchestrates the flow of data between entities and repositories.

**Components**:
- `auth/` - Authentication use cases
- `mongo/` - MongoDB operation use cases
- `etcd/` - Etcd operation use cases
- `consul/` - Consul operation use cases
- `cdn/` - CDN management use cases

**Example**:
```go
// internal/usecase/auth/auth_usecase.go
type AuthUseCase struct {
    userRepo  repository.UserRepository
    groupRepo repository.GroupRepository
}

func (uc *AuthUseCase) Login(ctx context.Context, username, password string) (*entity.User, error) {
    user, err := uc.userRepo.FindByUsername(ctx, username)
    if err != nil {
        return nil, ErrInvalidCredentials
    }
    // Business logic here...
    return user, nil
}
```

### 3. Adapter Layer (Interface Adapters)

**Location**: `internal/adapter/`

**Purpose**: Converts data between use cases and external systems. Contains HTTP handlers, presenters, and repository implementations.

**Components**:
- `handler/http/` - HTTP request handlers
- `handler/middleware/` - HTTP middleware (JWT auth)
- `presenter/` - Response formatters
- `repository/` - Repository implementations (optional, can be in infrastructure)

**Example**:
```go
// internal/adapter/handler/http/mongo_handler.go
type MongoHandler struct {
    useCase *mongo.MongoUseCase
}

func (h *MongoHandler) GetDocument(c *gin.Context) {
    var query entity.MongoQuery
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := h.useCase.GetDocuments(c.Request.Context(), &query)
    // Handle response...
}
```

### 4. Infrastructure Layer (Frameworks & Drivers)

**Location**: `internal/infrastructure/`

**Purpose**: Contains external frameworks, database implementations, and third-party integrations.

**Components**:
- `database/mongodb/` - MongoDB client and repository implementations
- `database/mysql/` - MySQL implementations
- `client/etcd/` - Etcd client implementation
- `client/consul/` - Consul client implementation
- `client/cloudflare/` - Cloudflare API client
- `config/` - Configuration management

**Example**:
```go
// internal/infrastructure/database/mongodb/mongo_repository.go
type MongoDBRepository struct {
    client *mongo.Client
}

func (r *MongoDBRepository) Find(ctx context.Context, database, collection string, filter bson.M, limit int64) ([]bson.M, error) {
    coll := r.client.Database(database).Collection(collection)
    // MongoDB-specific implementation...
}
```

## Directory Structure

```
golang-gin-sample/
├── internal/                          # Private application code
│   ├── domain/                        # Domain layer (innermost)
│   │   ├── entity/                    # Domain entities
│   │   │   ├── user.go                # User and Group entities
│   │   │   ├── document.go            # MongoDB document entities
│   │   │   ├── keyvalue.go            # Key-value entities (etcd, Consul)
│   │   │   └── cdn.go                 # CDN entities
│   │   └── repository/                # Repository interfaces (ports)
│   │       ├── user_repository.go     # User/Group repository interfaces
│   │       ├── mongo_repository.go    # MongoDB repository interface
│   │       ├── keyvalue_repository.go # Etcd/Consul repository interfaces
│   │       └── cdn_repository.go      # CDN repository interfaces
│   │
│   ├── usecase/                       # Use case layer
│   │   ├── auth/                      # Authentication use cases
│   │   │   └── auth_usecase.go
│   │   ├── mongo/                     # MongoDB use cases
│   │   │   └── mongo_usecase.go
│   │   ├── etcd/                      # Etcd use cases
│   │   │   └── etcd_usecase.go
│   │   ├── consul/                    # Consul use cases
│   │   │   └── consul_usecase.go
│   │   └── cdn/                       # CDN use cases
│   │       ├── cloudflare_usecase.go
│   │       ├── verizon_usecase.go
│   │       └── greycdn_usecase.go
│   │
│   ├── adapter/                       # Adapter layer
│   │   ├── handler/                   # HTTP handlers
│   │   │   ├── http/                  # HTTP request handlers
│   │   │   │   ├── mongo_handler.go
│   │   │   │   ├── etcd_handler.go
│   │   │   │   └── consul_handler.go
│   │   │   └── middleware/            # HTTP middleware
│   │   │       └── jwt.go
│   │   ├── presenter/                 # Response formatters
│   │   └── repository/                # Repository adapters
│   │
│   └── infrastructure/                # Infrastructure layer (outermost)
│       ├── database/                  # Database implementations
│       │   ├── mongodb/               # MongoDB implementation
│       │   │   ├── mongo_repository.go
│       │   │   ├── user_repository.go
│       │   │   └── group_repository.go
│       │   └── mysql/                 # MySQL implementation
│       ├── client/                    # External client implementations
│       │   ├── etcd/                  # Etcd client
│       │   │   └── etcd_repository.go
│       │   ├── consul/                # Consul client
│       │   │   └── consul_repository.go
│       │   ├── cloudflare/            # Cloudflare client
│       │   ├── verizon/               # Verizon client
│       │   └── greycdn/               # GreyCDN client
│       └── config/                    # Configuration management
│           └── config.go
│
├── cmd/                               # Application entry points
│   ├── mongo-api/
│   │   ├── main.go                    # Original implementation
│   │   └── main_clean.go              # Clean Architecture version
│   ├── etcd-api/
│   │   ├── main.go
│   │   └── main_clean.go
│   ├── consul-api/
│   │   ├── main.go
│   │   └── main_clean.go
│   └── cdn-api/
│       └── main.go
│
├── pkg/                               # Public shared libraries (legacy)
└── api/                               # OpenAPI/Swagger specs
```

## Dependency Flow

The Dependency Rule is fundamental to Clean Architecture:

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
│         Depends on: Domain entities & interfaces        │
└─────────────────────┬───────────────────────────────────┘
                      ↑
┌─────────────────────┴───────────────────────────────────┐
│                   ADAPTER LAYER                         │
│         (HTTP Handlers, Presenters)                     │
│         Depends on: Use cases & Domain                  │
└─────────────────────┬───────────────────────────────────┘
                      ↑
┌─────────────────────┴───────────────────────────────────┐
│                INFRASTRUCTURE LAYER                     │
│     (Database, External APIs, Frameworks)               │
│   Depends on: Implements Domain interfaces              │
└─────────────────────────────────────────────────────────┘
```

**Key Principles**:
- Inner layers know nothing about outer layers
- Dependencies point INWARD
- Domain layer has NO external dependencies
- Outer layers implement interfaces defined in inner layers

## Usage Guide

### Running the Clean Architecture Version

1. **Set environment variables**:
```bash
export MONGO_ADDRESS="localhost:27017"
export MONGO_USER="admin"
export MONGO_PASS="password"
export MONGO_DATABASE="test"
export JWT_KEY="your-secret-key"
export JWT_REALM="your-realm"
export JWT_TOKEN_HEADER="Bearer"
```

2. **Run mongo-api with Clean Architecture**:
```bash
go run cmd/mongo-api/main_clean.go
```

3. **Run etcd-api with Clean Architecture**:
```bash
export ETCD_ENDPOINTS="localhost:2379"
go run cmd/etcd-api/main_clean.go
```

4. **Run consul-api with Clean Architecture**:
```bash
export CONSUL_ADDRESS="localhost:8500"
go run cmd/consul-api/main_clean.go
```

### API Endpoints

All services maintain the same API endpoints as before:

**Authentication** (all services):
- `POST /auth/login` - Login and get JWT token
- `GET /auth/refresh_token` - Refresh JWT token
- `POST /auth/logout` - Logout

**MongoDB API**:
- `POST /api/v1/getDocument` - Query documents
- `POST /api/v1/insertSingleDocument` - Insert document
- `POST /api/v1/insertManyDocument` - Insert multiple documents
- `PUT /api/v1/updateSingleDocument` - Update document
- `DELETE /api/v1/deleteSingleDocument` - Delete document

**Etcd API**:
- `POST /api/v1/getKey` - Get key value
- `POST /api/v1/getKeyWithPrefix` - Get keys with prefix
- `POST /api/v1/putKey` - Store key-value
- `DELETE /api/v1/deleteKey` - Delete key

**Consul API**:
- `POST /api/v1/getKey` - Get key value
- `POST /api/v1/listKey` - List keys with prefix
- `POST /api/v1/putKey` - Store key-value
- `DELETE /api/v1/deleteKey` - Delete key

## Migration Guide

### Migrating Existing Code

To migrate existing code to Clean Architecture:

1. **Identify the layer**: Determine which layer your code belongs to
2. **Extract business logic**: Move business logic to use cases
3. **Define interfaces**: Create repository interfaces in domain layer
4. **Implement repositories**: Implement interfaces in infrastructure layer
5. **Update handlers**: Use use cases instead of direct database calls

### Example Migration

**Before** (old structure):
```go
// cmd/mongo-api/api/get.go
func GetDocument(c *gin.Context) {
    var reqBody MongoAPIGetBodyTop
    c.ShouldBindJSON(&reqBody)

    query := customMongoDatabase.MongoResultComplexQueryHelper{
        Cl:       mongoClient,
        Db:       customMongoSetup.MongoDatabase,
        Coll:     reqBody.Collection,
        FindData: reqBody.Body,
    }

    resp, err := query.FindDataComplexQueryFromMongo()
    c.JSON(http.StatusOK, gin.H{"result": resp})
}
```

**After** (Clean Architecture):
```go
// internal/adapter/handler/http/mongo_handler.go
func (h *MongoHandler) GetDocument(c *gin.Context) {
    var query entity.MongoQuery
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := h.useCase.GetDocuments(c.Request.Context(), &query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}
```

### Benefits of Migration

1. **Testability**: Business logic can be unit tested without HTTP or database
2. **Flexibility**: Easy to swap MongoDB for another database
3. **Maintainability**: Clear separation makes code easier to understand
4. **Reusability**: Use cases can be reused across different interfaces (HTTP, gRPC, CLI)

## Testing

### Unit Testing Use Cases

```go
func TestAuthUseCase_Login(t *testing.T) {
    // Create mock repositories
    mockUserRepo := &MockUserRepository{}
    mockGroupRepo := &MockGroupRepository{}

    // Create use case
    authUC := auth.NewAuthUseCase(mockUserRepo, mockGroupRepo)

    // Test login
    user, err := authUC.Login(context.Background(), "testuser", "password")

    // Assert results
    assert.NoError(t, err)
    assert.Equal(t, "testuser", user.Username)
}
```

### Integration Testing

Integration tests can test the full stack by injecting real repositories.

## Best Practices

1. **Keep domain layer pure**: No external dependencies in domain layer
2. **Use dependency injection**: Inject dependencies through constructors
3. **Define interfaces in inner layers**: Interfaces belong where they're used
4. **Return domain entities from use cases**: Don't leak infrastructure types
5. **Handle errors properly**: Use custom error types in domain layer
6. **Keep handlers thin**: Handlers should only handle HTTP concerns

## Next Steps

1. **Migrate CDN API**: Apply Clean Architecture to cdn-api
2. **Add integration tests**: Test the full stack with real dependencies
3. **Add metrics**: Instrument use cases for observability
4. **Add caching**: Implement caching at the repository level
5. **Generate OpenAPI specs**: Update Swagger docs for new structure

## References

- [The Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Clean Architecture in Go](https://github.com/bxcodec/go-clean-arch)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
