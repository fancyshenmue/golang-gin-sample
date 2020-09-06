package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "golang-gin-sample/api/consul-api/swagger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	jwt "github.com/appleboy/gin-jwt/v2"

	consulapi "golang-gin-sample/cmd/consul-api/api"
	consulAPIRouters "golang-gin-sample/cmd/consul-api/route"
)

// @title Consul RESTful API
// @version 1.0
// @description Consul RESTful API
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authMiddleware := consulapi.AuthMiddleware("consul-api")

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	consulAPIRouters.Consul(router)
	consulAPIRouters.ConsulAuth(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
