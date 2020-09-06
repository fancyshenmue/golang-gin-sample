package route

import (
	"github.com/gin-gonic/gin"

	cdnapi "golang-gin-sample/cmd/cdn-api/api"
)

func CDNAPIAuth(route *gin.Engine) {
	authMiddleware := cdnapi.AuthMiddleware("cdn-api")

	auth := route.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/logout", authMiddleware.LogoutHandler)
}
