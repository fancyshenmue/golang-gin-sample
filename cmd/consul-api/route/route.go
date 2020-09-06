package route

import (
	"github.com/gin-gonic/gin"

	consulapi "golang-gin-sample/cmd/consul-api/api"
)

func Consul(route *gin.Engine) {
	authMiddleware := consulapi.AuthMiddleware("consul-api")

	consul := route.Group("/api/v1")

	consul.Use(authMiddleware.MiddlewareFunc())

	// Query
	consul.POST("/getKey", consulapi.GetKey)

	// Insert
	consul.POST("/putKey", consulapi.PutKey)

	// Delete
	consul.DELETE("/deleteKey", consulapi.DeleteKey)
}

func ConsulAuth(route *gin.Engine) {
	authMiddleware := consulapi.AuthMiddleware("consul-api")

	auth := route.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/logout", authMiddleware.LogoutHandler)
}
