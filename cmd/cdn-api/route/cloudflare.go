package route

import (
	"github.com/gin-gonic/gin"

	cdnapi "github.com/fancyshenmue/golang-gin-sample/cmd/cdn-api/api"
)

func CDNAPICloudflare(route *gin.Engine) {
	authMiddleware := cdnapi.AuthMiddleware("cdn-api")

	/* account */
	cdnAccount := route.Group("/api/v1/cloudflare/account")
	cdnAccount.Use(authMiddleware.MiddlewareFunc())

	cdnAccount.POST("/list/accounts", cdnapi.CloudflareAccountListAccounts)
	cdnAccount.POST("/details", cdnapi.CloudflareAccountDetails)
	/* end of account */

	/* zone */
	cdnZone := route.Group("/api/v1/cloudflare/zone")
	cdnZone.Use(authMiddleware.MiddlewareFunc())
	// cdnZone.Use(cdnapi.CloudflareAccountDetailsMiddleware())

	cdnZone.POST("/list/zones", cdnapi.CloudflareZoneListZones)
	cdnZone.POST("/purge/all/files", cdnapi.CloudflareZonePurgeAllFiles)
	cdnZone.POST("/purge/files/by/url", cdnapi.CloudflareZonePurgeFilesByURL)
	/* end of zone */
}
