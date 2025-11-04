package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	httpHandler "github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/http"
	"github.com/fancyshenmue/golang-gin-sample/internal/adapter/handler/middleware"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/client/consul"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/client/etcd"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/config"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/database/mongodb"
	authUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/auth"
	consulUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/consul"
	etcdUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/etcd"
	mongoUseCase "github.com/fancyshenmue/golang-gin-sample/internal/usecase/mongo"
)

// @title Unified RESTful API
// @version 2.0
// @description Unified RESTful API with Clean Architecture - MongoDB, Etcd, Consul, CDN
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	log.Println("Starting Unified API Server with Clean Architecture...")

	// Initialize MongoDB connection
	mongoClient, err := mongodb.Connect(cfg.MongoDB.URI, &cfg.MongoDB.Credential)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())
	log.Println("✓ Connected to MongoDB")

	// Initialize Etcd connection (optional - will log warning if not available)
	var etcdRepo *etcd.EtcdClientRepository
	if len(cfg.Etcd.Endpoints) > 0 && cfg.Etcd.Endpoints[0] != "" {
		etcdRepo, err = etcd.NewEtcdClientRepository(
			cfg.Etcd.Endpoints,
			os.Getenv("ETCD_CONNECT_USER"),
			os.Getenv("ETCD_CONNECT_PASSWD"),
		)
		if err != nil {
			log.Printf("⚠ Warning: Failed to connect to Etcd: %v (Etcd endpoints will be unavailable)", err)
		} else {
			defer etcdRepo.Close()
			log.Println("✓ Connected to Etcd")
		}
	} else {
		log.Println("⚠ Etcd not configured (skipping)")
	}

	// Initialize Consul connection (optional - will log warning if not available)
	var consulRepo *consul.ConsulClientRepository
	if cfg.Consul.Address != "" {
		consulRepo, err = consul.NewConsulClientRepository(
			cfg.Consul.Address,
			os.Getenv("CONSUL_TOKEN"),
		)
		if err != nil {
			log.Printf("⚠ Warning: Failed to connect to Consul: %v (Consul endpoints will be unavailable)", err)
		} else {
			log.Println("✓ Connected to Consul")
		}
	} else {
		log.Println("⚠ Consul not configured (skipping)")
	}

	// Initialize repositories
	mongoRepo := mongodb.NewMongoDBRepository(mongoClient)
	userRepo := mongodb.NewUserMongoRepository(mongoClient, cfg.MongoDB.Database)
	groupRepo := mongodb.NewGroupMongoRepository(mongoClient, cfg.MongoDB.Database)

	// Initialize use cases
	authUC := authUseCase.NewAuthUseCase(userRepo, groupRepo)
	mongoUC := mongoUseCase.NewMongoUseCase(mongoRepo)

	var etcdUC *etcdUseCase.EtcdUseCase
	if etcdRepo != nil {
		etcdUC = etcdUseCase.NewEtcdUseCase(etcdRepo)
	}

	var consulUC *consulUseCase.ConsulUseCase
	if consulRepo != nil {
		consulUC = consulUseCase.NewConsulUseCase(consulRepo)
	}

	// Initialize handlers
	mongoHandler := httpHandler.NewMongoHandler(mongoUC)

	var etcdHandler *httpHandler.EtcdHandler
	if etcdUC != nil {
		etcdHandler = httpHandler.NewEtcdHandler(etcdUC)
	}

	var consulHandler *httpHandler.ConsulHandler
	if consulUC != nil {
		consulHandler = httpHandler.NewConsulHandler(consulUC)
	}

	// Initialize JWT middleware (using "api" as the required group)
	jwtMiddleware := middleware.NewJWTMiddleware(cfg, authUC, "api")

	// Setup Gin router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Unified API - Clean Architecture",
			"version": "2.0",
			"services": gin.H{
				"mongodb": "available",
				"etcd":    getServiceStatus(etcdHandler),
				"consul":  getServiceStatus(consulHandler),
			},
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"services": gin.H{
				"mongodb": "connected",
				"etcd":    getServiceStatus(etcdHandler),
				"consul":  getServiceStatus(consulHandler),
			},
		})
	})

	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", jwtMiddleware.LoginHandler)
		auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
		auth.POST("/logout", jwtMiddleware.LogoutHandler)
	}

	// Protected API routes
	api := router.Group("/api/v1")
	api.Use(jwtMiddleware.MiddlewareFunc())
	{
		// MongoDB endpoints
		mongo := api.Group("/mongo")
		{
			mongo.POST("/getDocument", mongoHandler.GetDocument)
			mongo.POST("/insertSingleDocument", mongoHandler.InsertSingleDocument)
			mongo.POST("/insertManyDocument", mongoHandler.InsertManyDocument)
			mongo.PUT("/updateSingleDocument", mongoHandler.UpdateSingleDocument)
			mongo.DELETE("/deleteSingleDocument", mongoHandler.DeleteSingleDocument)
		}

		// Etcd endpoints (if available)
		if etcdHandler != nil {
			etcdGroup := api.Group("/etcd")
			{
				etcdGroup.POST("/getKey", etcdHandler.GetKey)
				etcdGroup.POST("/getKeyWithPrefix", etcdHandler.GetKeyWithPrefix)
				etcdGroup.POST("/putKey", etcdHandler.PutKey)
				etcdGroup.DELETE("/deleteKey", etcdHandler.DeleteKey)
			}
		}

		// Consul endpoints (if available)
		if consulHandler != nil {
			consulGroup := api.Group("/consul")
			{
				consulGroup.POST("/getKey", consulHandler.GetKey)
				consulGroup.POST("/listKey", consulHandler.ListKeys)
				consulGroup.POST("/putKey", consulHandler.PutKey)
				consulGroup.DELETE("/deleteKey", consulHandler.DeleteKey)
			}
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	port := cfg.Server.Port
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📚 API Documentation: http://localhost:%s/swagger/index.html", port)
	log.Printf("🏥 Health Check: http://localhost:%s/health", port)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getServiceStatus returns the status of a service handler
func getServiceStatus(handler interface{}) string {
	if handler == nil {
		return "not configured"
	}
	return "available"
}
