package app

import (
	"github.com/gin-gonic/gin"

	customMongoAPI "golang-gin-sample/cmd/app/mongo-api/app"
)

func Route(route *gin.Engine) {
	// authMiddleware := mongoapi.AuthMiddleware("mongo-api")

	mongo := route.Group("/api/v1")

	// mongo.Use(authMiddleware.MiddlewareFunc())

	// Query
	mongo.POST("/getDocument", customMongoAPI.FindDocument)
	// mongo.POST("/getDocumentRegex", mongoapi.GetDocumentRegex)

	// Insert
	mongo.POST("/insertDocument", customMongoAPI.InsertDocument)
	// mongo.POST("/insertManyDocument", mongoapi.InsertManyDocument)

	// Update
	mongo.PUT("/updateDocument", customMongoAPI.UpdateDocument)

	// DELETE
	mongo.DELETE("/deleteDocument", customMongoAPI.DeleteDocument)
}

// func MongoAuth(route *gin.Engine) {
// 	authMiddleware := mongoapi.AuthMiddleware("mongo-api")

// 	auth := route.Group("/auth")
// 	auth.POST("/login", authMiddleware.LoginHandler)
// 	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// 	auth.POST("/logout", authMiddleware.LogoutHandler)
// }
