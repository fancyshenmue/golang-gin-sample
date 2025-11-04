package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	httpHandler "github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/http"
	"github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/middleware"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/client/etcd"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/config"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/database/mongodb"
	authUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/auth"
	etcdUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/etcd"
)

// @title Etcd RESTful API
// @version 2.0
// @description Etcd RESTful API with Clean Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize etcd connection
	etcdRepo, err := etcd.NewEtcdClientRepository(
		cfg.Etcd.Endpoints,
		os.Getenv("ETCD_CONNECT_USER"),
		os.Getenv("ETCD_CONNECT_PASSWD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdRepo.Close()

	log.Println("Successfully connected to etcd")

	// Initialize MongoDB for authentication
	mongoClient, err := mongodb.Connect(cfg.MongoDB.URI, &cfg.MongoDB.Credential)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userRepo := mongodb.NewUserMongoRepository(mongoClient, cfg.MongoDB.Database)
	groupRepo := mongodb.NewGroupMongoRepository(mongoClient, cfg.MongoDB.Database)

	// Initialize use cases
	authUC := authUseCase.NewAuthUseCase(userRepo, groupRepo)
	etcdUC := etcdUseCase.NewEtcdUseCase(etcdRepo)

	// Initialize handlers
	etcdHandler := httpHandler.NewEtcdHandler(etcdUC)

	// Initialize JWT middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg, authUC, "etcd-api")

	// Setup router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", jwtMiddleware.LoginHandler)
		auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
		auth.POST("/logout", jwtMiddleware.LogoutHandler)
	}

	// API routes (protected)
	api := router.Group("/api/v1")
	api.Use(jwtMiddleware.MiddlewareFunc())
	{
		api.POST("/getKey", etcdHandler.GetKey)
		api.POST("/getKeyWithPrefix", etcdHandler.GetKeyWithPrefix)
		api.POST("/putKey", etcdHandler.PutKey)
		api.DELETE("/deleteKey", etcdHandler.DeleteKey)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Start server
	port := cfg.Server.Port
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
