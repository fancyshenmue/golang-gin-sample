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
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/client/consul"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/config"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/database/mongodb"
	authUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/auth"
	consulUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/consul"
)

// @title Consul RESTful API
// @version 2.0
// @description Consul RESTful API with Clean Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Consul connection
	consulRepo, err := consul.NewConsulClientRepository(
		cfg.Consul.Address,
		os.Getenv("CONSUL_TOKEN"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Consul: %v", err)
	}

	log.Println("Successfully connected to Consul")

	// Initialize MongoDB for authentication
	mongoClient, err := mongodb.Connect(cfg.MongoDB.URI, &cfg.MongoDB.Credential)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userRepo := mongodb.NewUserMongoRepository(mongoClient, cfg.MongoDB.Database)
	groupRepo := mongodb.NewGroupMongoRepository(mongoClient, cfg.MongoDB.Database)

	// Initialize use cases
	authUC := authUseCase.NewAuthUseCase(userRepo, groupRepo)
	consulUC := consulUseCase.NewConsulUseCase(consulRepo)

	// Initialize handlers
	consulHandler := httpHandler.NewConsulHandler(consulUC)

	// Initialize JWT middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg, authUC, "consul-api")

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
		api.POST("/getKey", consulHandler.GetKey)
		api.POST("/listKey", consulHandler.ListKeys)
		api.POST("/putKey", consulHandler.PutKey)
		api.DELETE("/deleteKey", consulHandler.DeleteKey)
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
