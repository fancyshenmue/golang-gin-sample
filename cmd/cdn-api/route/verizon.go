package route

import (
	"github.com/gin-gonic/gin"

	cdnapi "golang-gin-sample/cmd/cdn-api/api"
)

func CDNAPIVerizon(route *gin.Engine) {
	authMiddleware := cdnapi.AuthMiddleware("cdn-api")

	/* cache */

	cdnEdgeCache := route.Group("/api/v1/verizon/edge/cache")
	cdnEdgeCache.Use(authMiddleware.MiddlewareFunc())

	cdnEdgeCache.PUT("/purge/content", cdnapi.VerizonEdgeCachePurgeContent)

	/* end of cache */

	/* customer oringins */

	cdnEdgeCustomerOringins := route.Group("/api/v1/verizon/edge/customer/origins")
	cdnEdgeCustomerOringins.Use(authMiddleware.MiddlewareFunc())

	cdnEdgeCustomerOringins.GET("/get/:originid", cdnapi.VerizonEdgeCustomerOriginGet)
	cdnEdgeCustomerOringins.POST("/add/httpLarge", cdnapi.VerizonEdgeCustomerOriginAddHTTPLarge)
	cdnEdgeCustomerOringins.DELETE("/delete/:originid", cdnapi.VerizonEdgeCustomerOriginDelete)

	/* end of customer oringins */

	/* edge cnames */

	cdnEdgeCNAME := route.Group("/api/v1/verizon/edge/cnames")
	cdnEdgeCNAME.Use(authMiddleware.MiddlewareFunc())

	cdnEdgeCNAME.GET("/list/httpLarge", cdnapi.VerizonEdgeCNAMEListHTTPLarge)
	cdnEdgeCNAME.GET("/get/:cnameid", cdnapi.VerizonEdgeCNAMEGet)
	cdnEdgeCNAME.POST("/add", cdnapi.VerizonEdgeCNAMEAdd)
	cdnEdgeCNAME.DELETE("/delete/:cnameid", cdnapi.VerizonEdgeCNAMEDelete)

	/* end of edge cnames */
}
