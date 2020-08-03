package route

import (
	"github.com/gin-gonic/gin"

	mongoapi "github.com/fancyshenmue/golang-gin-sample/cmd/mongo-api/api"
)

func MongoRoute(route *gin.Engine) {
	authMiddleware := mongoapi.AuthMiddleware("mongo-api")

	mongo := route.Group("/api/v1")

	mongo.Use(authMiddleware.MiddlewareFunc())

	// Query
	mongo.POST("/getDocument", mongoapi.GetDocument)
	mongo.POST("/getDocumentRegex", mongoapi.GetDocumentRegex)

	// Insert
	mongo.POST("/insertSingleDocument/:collection", mongoapi.InsertSingleDocument)
	mongo.POST("/insertManyDocument/:collection", mongoapi.InsertManyDocument)

	// Update
	mongo.PUT("/updateSingleDocument/:collection", mongoapi.UpdateSingleDocument)

	// DELETE
	mongo.DELETE("/deleteSingleDocument/:collection", mongoapi.DeleteSingleDocument)
}

func MongoAuth(route *gin.Engine) {
	authMiddleware := mongoapi.AuthMiddleware("mongo-api")

	auth := route.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/logout", authMiddleware.LogoutHandler)
}
