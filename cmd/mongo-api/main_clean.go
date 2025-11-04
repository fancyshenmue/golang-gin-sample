package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/fancyshenmue/golang-gin-sample/api/mongo-api/swagger/docs"

	httpHandler "github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/http"
	"github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/middleware"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/config"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/database/mongodb"
	authUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/auth"
	mongoUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/mongo"
)

// @title Mongo RESTful API
// @version 2.0
// @description Mongo RESTful API with Clean Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB connection
	mongoClient, err := mongodb.Connect(cfg.MongoDB.URI, &cfg.MongoDB.Credential)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	log.Println("Successfully connected to MongoDB")

	// Initialize repositories
	mongoRepo := mongodb.NewMongoDBRepository(mongoClient)
	userRepo := mongodb.NewUserMongoRepository(mongoClient, cfg.MongoDB.Database)
	groupRepo := mongodb.NewGroupMongoRepository(mongoClient, cfg.MongoDB.Database)

	// Initialize use cases
	authUC := authUseCase.NewAuthUseCase(userRepo, groupRepo)
	mongoUC := mongoUseCase.NewMongoUseCase(mongoRepo)

	// Initialize handlers
	mongoHandler := httpHandler.NewMongoHandler(mongoUC)

	// Initialize JWT middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg, authUC, "mongo-api")

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
		// MongoDB operations
		api.POST("/getDocument", mongoHandler.GetDocument)
		api.POST("/insertSingleDocument", mongoHandler.InsertSingleDocument)
		api.POST("/insertManyDocument", mongoHandler.InsertManyDocument)
		api.PUT("/updateSingleDocument", mongoHandler.UpdateSingleDocument)
		api.DELETE("/deleteSingleDocument", mongoHandler.DeleteSingleDocument)
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
