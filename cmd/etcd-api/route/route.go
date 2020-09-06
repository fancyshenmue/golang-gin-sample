package route

import (
	"os"

	"github.com/gin-gonic/gin"

	etcdapi "golang-gin-sample/cmd/etcd-api/api"
)

func Etcd(route *gin.Engine) {
	authMiddleware := etcdapi.AuthMiddleware("etcd-api")

	env := etcdapi.Environment{
		os.Getenv("ENV"),
	}

	etcd := route.Group("/api/v1")

	etcd.Use(authMiddleware.MiddlewareFunc())

	// Query
	etcd.POST("/getKey", env.GetKey)
	etcd.POST("/getKeyWithPrefix", env.GetKeyWithPrefix)

	// Insert
	etcd.POST("/putKey", env.PutKey)

	// Delete
	etcd.DELETE("/deleteKey", env.DeleteKey)
}

func EtcdAuth(route *gin.Engine) {
	authMiddleware := etcdapi.AuthMiddleware("etcd-api")

	auth := route.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/logout", authMiddleware.LogoutHandler)
}
