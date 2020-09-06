package route

import (
	"github.com/gin-gonic/gin"

	cdnapi "golang-gin-sample/cmd/cdn-api/api"
)

func CDNAPIGreyCDN(route *gin.Engine) {
	authMiddleware := cdnapi.AuthMiddleware("cdn-api")

	/* site */
	cdnSite := route.Group("/api/v1/greycdn/site")
	cdnSite.Use(authMiddleware.MiddlewareFunc())

	cdnSite.GET("/list/all", cdnapi.GreyCDNSiteListAll)
	/* end of site */

	/* cache */
	cdnCache := route.Group("/api/v1/greycdn/cache")
	cdnCache.Use(authMiddleware.MiddlewareFunc())

	cdnCache.POST("/purge/by-site", cdnapi.GreyCDNCachePurgeBySite)
	cdnCache.POST("/purge/by-site-uri", cdnapi.GreyCDNCachePurgeBySiteURI)
	/* end of cache */
}
