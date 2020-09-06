package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "golang-gin-sample/api/cdn-api/swagger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	jwt "github.com/appleboy/gin-jwt/v2"

	CDNAPI "golang-gin-sample/cmd/cdn-api/api"
	CDNAPIRouters "golang-gin-sample/cmd/cdn-api/route"
)

// @title CDN RESTful API
// @version 1.0
// @description CDN RESTful API
// @termsOfService http://swagger.io/terms/

// @contact.name Charles Hsu
// @contact.email fancyshenmue@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	authMiddleware := CDNAPI.AuthMiddleware("cdn-api")

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	CDNAPIRouters.CDNAPIAuth(router)
	CDNAPIRouters.CDNAPICloudflare(router)
	CDNAPIRouters.CDNAPIGreyCDN(router)
	CDNAPIRouters.CDNAPIVerizon(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
