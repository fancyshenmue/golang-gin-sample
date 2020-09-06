package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	customHttp "golang-gin-sample/pkg/http"
	customLogHandle "golang-gin-sample/pkg/loghandle"

	"github.com/gin-gonic/gin"
)

// VerizonEdgeCNAMEListHTTPLarge | Verizon Edge CNAME List HTTP Large
// @Summary Verizon Edge CNAME List HTTP Large
// @Id 1
// @Tags Verizon CDN (CNAME)
// @version 1.0
// @accept  json
// @produce json
// @Success 200 string string
// @Router /api/v1/verizon/edge/cnames/list/httpLarge [get]
func VerizonEdgeCNAMEListHTTPLarge(c *gin.Context) {
	var (
		resBody []VerizoneCDNCnameHttpLargeListResponse
	)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["edgeCNAME"]["getAllHTTPLarge"]["method"],
		URL:           fmt.Sprintf("%s", verizoneCDNAPIURI["edgeCNAME"]["getAllHTTPLarge"]["url"]),
		RequestHeader: verizonCDNTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCNAMEListHTTPLarge",
			"VerizonEdgeCNAMEListHTTPLarge",
			"VerizonEdgeCNAMEListHTTPLarge (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCNAMEGet | Verizon Edge CNAME Get
// @Summary Verizon Edge CNAME Get
// @Id 1
// @Tags Verizon CDN (CNAME)
// @version 1.0
// @accept  json
// @produce json
// @param cnameid path string true "cnameid"
// @Success 200 string string
// @Router /api/v1/verizon/edge/cnames/get/{cnameid} [get]
func VerizonEdgeCNAMEGet(c *gin.Context) {
	var (
		resBody VerizoneCDNCnameHttpLargeListResponse
	)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["edgeCNAME"]["getEdgeCNAME"]["method"],
		URL:           fmt.Sprintf("%s/%s", verizoneCDNAPIURI["edgeCNAME"]["getEdgeCNAME"]["url"], c.Param("cnameid")),
		RequestHeader: verizonCDNTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCNAMEGet",
			"VerizonEdgeCNAMEGet",
			"VerizonEdgeCNAMEGet (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCNAMEAdd | Verizon Edge CNAME Add
// @Summary Verizon Edge CNAME Add
// @Id 1
// @Tags Verizon CDN (CNAME)
// @version 1.0
// @accept  json
// @produce json
// @param params body VerizoneCDNCnameAddEdgeCNAMEsRequest true "body"
// @Success 200 string string
// @Router /api/v1/verizon/edge/cnames/add [post]
func VerizonEdgeCNAMEAdd(c *gin.Context) {
	var (
		reqBody VerizoneCDNCnameAddEdgeCNAMEsRequest
		resBody VerizoneCDNResponseBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// set default media type http large
	if reqBody.MediaTypeID == 0 {
		reqBody.MediaTypeID = verizonMediaTypeIDHTTPLarge
	}

	// set default origin ID
	if reqBody.OriginID == 0 {
		reqBody.OriginID = -1
	}

	reqBodyData := new(bytes.Buffer)
	json.NewEncoder(reqBodyData).Encode(reqBody)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["edgeCNAME"]["addEdgeCNAME"]["method"],
		URL:           fmt.Sprintf("%s", verizoneCDNAPIURI["edgeCNAME"]["addEdgeCNAME"]["url"]),
		RequestHeader: verizonCDNCommonHeader,
		RequestBody:   reqBodyData,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCNAMEAdd",
			"VerizonEdgeCNAMEAdd",
			"VerizonEdgeCNAMEAdd (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCNAMEDelete | Verizon Edge CNAME Delete
// @Summary Verizon Edge CNAME Delete
// @Id 1
// @Tags Verizon CDN (CNAME)
// @version 1.0
// @accept  json
// @produce json
// @param cnameid path string true "cnameid"
// @Success 200 string string
// @Router /api/v1/verizon/edge/cnames/delete/{cnameid} [delete]
func VerizonEdgeCNAMEDelete(c *gin.Context) {
	var (
		resBody VerizoneCDNResponseBody
	)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["edgeCNAME"]["deleteEdgeCNAME"]["method"],
		URL:           fmt.Sprintf("%s/%s", verizoneCDNAPIURI["edgeCNAME"]["deleteEdgeCNAME"]["url"], c.Param("cnameid")),
		RequestHeader: verizonCDNTokenHeaderDELETE,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCNAMEDelete",
			"VerizonEdgeCNAMEDelete",
			"VerizonEdgeCNAMEDelete (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCustomerOriginGet | Verizon Edge Customer Origin Get
// @Summary Verizon Edge Customer Origin Get
// @Id 1
// @Tags Verizon Customer Oringins
// @version 1.0
// @accept  json
// @produce json
// @param originid path string true "originid"
// @Success 200 string string
// @Router /api/v1/verizon/edge/customer/origins/get/{originid} [get]
func VerizonEdgeCustomerOriginGet(c *gin.Context) {
	var (
		resBody VerizoneCDNResponseBody
	)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["customerOrigins"]["getHTTPLarge"]["method"],
		URL:           fmt.Sprintf("%s/%s", verizoneCDNAPIURI["customerOrigins"]["getHTTPLarge"]["url"], c.Param("originid")),
		RequestHeader: verizonCDNCommonHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCustomerOriginGet",
			"VerizonEdgeCustomerOriginGet",
			"VerizonEdgeCustomerOriginGet (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCustomerOriginAddHTTPLarge | Verizon Edge Customer Origin Add HTTP Large
// @Summary Verizon Edge Customer Origin Add HTTP Large
// @Id 1
// @Tags Verizon Customer Oringins
// @version 1.0
// @accept  json
// @produce json
// @param params body VerizoneCDNCustomerOriginAddHTTPLargeRequest true "body"
// @Success 200 string string
// @Router /api/v1/verizon/edge/customer/origins/add/httpLarge [post]
func VerizonEdgeCustomerOriginAddHTTPLarge(c *gin.Context) {
	var (
		reqBody VerizoneCDNCustomerOriginAddHTTPLargeRequest
		resBody VerizoneCDNResponseBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reqBodyData := new(bytes.Buffer)
	json.NewEncoder(reqBodyData).Encode(reqBody)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["customerOrigins"]["addHTTPLarge"]["method"],
		URL:           fmt.Sprintf("%s", verizoneCDNAPIURI["customerOrigins"]["addHTTPLarge"]["url"]),
		RequestHeader: verizonCDNCommonHeader,
		RequestBody:   reqBodyData,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCustomerOriginAdd",
			"VerizonEdgeCustomerOriginAdd",
			"VerizonEdgeCustomerOriginAdd (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCustomerOriginDelete | Verizon Edge Customer Origin Delete
// @Summary Verizon Edge Customer Origin Delete
// @Id 1
// @Tags Verizon Customer Oringins
// @version 1.0
// @accept  json
// @produce json
// @param originid path string true "originid"
// @Success 200 string string
// @Router /api/v1/verizon/edge/customer/origins/delete/{originid} [delete]
func VerizonEdgeCustomerOriginDelete(c *gin.Context) {
	var (
		resBody VerizoneCDNResponseBody
	)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["customerOrigins"]["delete"]["method"],
		URL:           fmt.Sprintf("%s/%s", verizoneCDNAPIURI["customerOrigins"]["delete"]["url"], c.Param("originid")),
		RequestHeader: verizonCDNCommonHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCustomerOriginDelete",
			"VerizonEdgeCustomerOriginDelete",
			"VerizonEdgeCustomerOriginDelete (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// VerizonEdgeCachePurgeContent | Verizon Edge Cache Purge Content
// @Summary Verizon Edge Cache Purge Content
// @Id 1
// @Tags Verizon Cache
// @version 1.0
// @accept  json
// @produce json
// @param params body VerizoneCDNCachePurgeContentRequest true "body"
// @Success 200 string string
// @Router /api/v1/verizon/edge/cache/purge/content [put]
func VerizonEdgeCachePurgeContent(c *gin.Context) {
	var (
		reqBody VerizoneCDNCachePurgeContentRequest
		resBody VerizoneCDNResponseBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// set default media type http large
	if reqBody.MediaType == 0 {
		reqBody.MediaType = verizonMediaTypeIDHTTPLarge
	}

	reqBodyData := new(bytes.Buffer)
	json.NewEncoder(reqBodyData).Encode(reqBody)

	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["cache"]["purgeContent"]["method"],
		URL:           fmt.Sprintf("%s", verizoneCDNAPIURI["cache"]["purgeContent"]["url"]),
		RequestHeader: verizonCDNCommonHeader,
		RequestBody:   reqBodyData,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCachePurgeContent",
			"VerizonEdgeCachePurgeContent",
			"VerizonEdgeCachePurgeContent (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}
