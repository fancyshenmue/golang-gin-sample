package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	customHttp "golang-gin-sample/pkg/http"
	customLogHandle "golang-gin-sample/pkg/loghandle"
)

// GreyCDNSiteListAll | Grey CDN Site List All
// @Summary Grey CDN Site List All
// @Id 1
// @Tags GreyCDN (Site Info)
// @version 1.0
// @accept  json
// @produce json
// @Success 200 string string
// @Router /api/v1/greycdn/site/list/all [get]
func GreyCDNSiteListAll(c *gin.Context) {
	var greySiteAll GreyCDNSiteListResponseTop

	req := &customHttp.HttpRequestHelper{
		Method:        greyCDNAPIURI["site"]["listAll"]["method"],
		URL:           greyCDNAPIURI["site"]["listAll"]["url"],
		RequestHeader: greyCDNTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &greySiteAll)
	if err != nil {
		customLogHandle.ErrorHandle(
			"GreyCDNSiteListAll",
			"GreyCDNSiteListAll",
			"GreyCDNSiteListAll (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, greySiteAll.Result)
}
