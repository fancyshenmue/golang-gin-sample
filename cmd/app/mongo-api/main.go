package main

import (
	"fmt"
	APIRouters "golang-gin-sample/cmd/app/mongo-api/route"
	"os"

	"github.com/gin-gonic/gin"

	_ "golang-gin-sample/api/app/mongo-api/swagger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	APIRouters.Route(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if os.Getenv("APP_PORT") == "" {
		router.Run(":8080")
	} else if os.Getenv("APP_PORT") != "" {
		router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	}
}
