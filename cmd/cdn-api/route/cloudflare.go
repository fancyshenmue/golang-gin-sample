package route

import (
	"github.com/gin-gonic/gin"

	cdnapi "golang-gin-sample/cmd/cdn-api/api"
)

func CDNAPICloudflare(route *gin.Engine) {
	authMiddleware := cdnapi.AuthMiddleware("cdn-api")

	/* account */
	cdnAccount := route.Group("/api/v1/cloudflare/account")
	cdnAccount.Use(authMiddleware.MiddlewareFunc())

	cdnAccount.POST("/list/accounts", cdnapi.CloudflareAccountListAccounts)
	cdnAccount.POST("/details", cdnapi.CloudflareAccountDetails)
	/* end of account */

	/* dns */
	cdnDNS := route.Group("/api/v1/cloudflare/dns")
	cdnDNS.Use(authMiddleware.MiddlewareFunc())

	cdnDNS.POST("/list/records", cdnapi.CloudflareDNSListDNSRecords)
	cdnDNS.POST("/create/record", cdnapi.CloudflareDNSCreateRecord)
	cdnDNS.POST("/record/detail", cdnapi.CloudflareDNSRecordDetail)
	cdnDNS.POST("/delete/record", cdnapi.CloudflareDNSDeleteDNSRecord)

	/* end of dns */

	/* zone */
	cdnZone := route.Group("/api/v1/cloudflare/zone")
	cdnZone.Use(authMiddleware.MiddlewareFunc())

	cdnZone.POST("/list/zones", cdnapi.CloudflareZoneListZones)
	cdnZone.POST("/create", cdnapi.CloudflareZoneCreateZone)
	cdnZone.POST("/details", cdnapi.CloudflareZoneZoneDetail)
	cdnZone.POST("/delete", cdnapi.CloudflareZoneDeleteZone)
	cdnZone.POST("/purge/all/files", cdnapi.CloudflareZonePurgeAllFiles)
	cdnZone.POST("/purge/files/by/url", cdnapi.CloudflareZonePurgeFilesByURL)
	/* end of zone */
}
