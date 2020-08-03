package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	customHttp "github.com/fancyshenmue/golang-gin-sample/pkg/http"
	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"github.com/gin-gonic/gin"
)

// GreyCDNCachePurgeBySite | Grey CDN Cache Purge By Site
// @Summary Grey CDN Cache Purge By Site
// @Id 1
// @Tags GreyCDN (Cache)
// @version 1.0
// @accept  json
// @produce json
// @param params body GreyCDNCacheBySiteRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/greycdn/cache/purge/by-site [post]
func GreyCDNCachePurgeBySite(c *gin.Context) {
	var (
		reqBody GreyCDNCacheBySiteRequestBody
		resBody GreyCDNCacheBySiteResponseBody
	)

	c.ShouldBindJSON(&reqBody)

	req := &customHttp.HttpRequestHelper{
		Method:        greyCDNAPIURI["cache"]["purgeBySite"]["method"],
		URL:           fmt.Sprintf("%s?uid=%s", greyCDNAPIURI["cache"]["purgeBySite"]["url"], reqBody.UID),
		RequestHeader: greyCDNTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"GreyCDNCachePurgeBySite",
			"GreyCDNCachePurgeBySite",
			"GreyCDNCachePurgeBySite (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// GreyCDNCachePurgeBySiteURI | Grey CDN Cache Purge By Site URI
// @Summary Grey CDN Cache Purge By Site URI
// @Id 1
// @Tags GreyCDN (Cache)
// @version 1.0
// @accept  json
// @produce json
// @param params body GreyCDNCacheBySiteURIRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/greycdn/cache/purge/by-site-uri [post]
func GreyCDNCachePurgeBySiteURI(c *gin.Context) {
	var (
		reqBody GreyCDNCacheBySiteURIRequestBody
		resBody GreyCDNCacheBySiteURIResponseBody
	)

	c.ShouldBindJSON(&reqBody)

	var reqBodyData = fmt.Sprintf(`{
	"uri":"%s"
}`, reqBody.URI)

	var reqBodyDataByte = []byte(reqBodyData)

	req := &customHttp.HttpRequestHelper{
		Method:        greyCDNAPIURI["cache"]["purgeBySiteURI"]["method"],
		URL:           fmt.Sprintf("%s?uid=%s", greyCDNAPIURI["cache"]["purgeBySiteURI"]["url"], reqBody.UID),
		RequestHeader: greyCDNTokenHeader,
		RequestBody:   bytes.NewBuffer(reqBodyDataByte),
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"GreyCDNCachePurgeBySiteURI",
			"GreyCDNCachePurgeBySiteURI",
			"GreyCDNCachePurgeBySiteURI (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}
