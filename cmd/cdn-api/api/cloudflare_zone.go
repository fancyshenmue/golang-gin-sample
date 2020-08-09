package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	customHttp "github.com/fancyshenmue/golang-gin-sample/pkg/http"
	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/publicsuffix"
)

// CloudflareZoneListZones | Cloudflare Zone List Zones
// @Summary Cloudflare Zone List Zones
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/list/zones [post]
func CloudflareZoneListZones(c *gin.Context) {
	api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	resp, err := api.ListZones()

	c.JSON(http.StatusOK, resp)
}

// CloudflareZoneCreateZone | Cloudflare Zone Create Zone
// @Summary Cloudflare Zone Create Zone
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareZoneCreateZoneRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/create [post]
func CloudflareZoneCreateZone(c *gin.Context) {
	var (
		reqBody              CloudflareZoneCreateZoneRequestBody
		resBody              CloudflareCDNResponseBody
		cloudflarePagination cloudflare.PaginationOptions
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   reqBody.AuthKey,
		"X-Auth-Email": reqBody.User,
	}

	// get account id
	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	cloudflareaccountInfo, _, _ := api.Accounts(cloudflarePagination)
	cloudflareaccountID := cloudflareaccountInfo[0].ID

	// default type full
	if reqBody.Type == "" {
		reqBody.Type = "full"
	}

	var (
		reqBodyData = fmt.Sprintf(`{
	"name": "%s",
	"account": {
		"id": "%s"
	},
	"jump_start": %t,
	"type": "%s"
}`, reqBody.Name, cloudflareaccountID, reqBody.JumpStart, reqBody.Type)

		reqBodyDataByte = []byte(reqBodyData)
	)

	req := &customHttp.HttpRequestHelper{
		Method:        cloudflareCDNAPIURI["zone"]["createZone"]["method"],
		URL:           fmt.Sprintf("%s", cloudflareCDNAPIURI["zone"]["createZone"]["url"]),
		RequestHeader: cloudflareTokenHeader,
		RequestBody:   bytes.NewBuffer(reqBodyDataByte),
	}

	resp := req.HttpRequest()

	err = json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"CloudflareZoneCreateZone",
			"CloudflareZoneCreateZone",
			"CloudflareZoneCreateZone (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// CloudflareZoneZoneDetail | Cloudflare Zone Zone Detail
// @Summary Cloudflare Zone Zone Detail
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareZoneZoneDetailRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/details [post]
func CloudflareZoneZoneDetail(c *gin.Context) {
	var (
		reqBody CloudflareZoneZoneDetailRequestBody
		resBody CloudflareCDNResponseBody
		zoneID  string
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   reqBody.AuthKey,
		"X-Auth-Email": reqBody.User,
	}

	// get zone list
	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	respZoneList, err := api.ListZones()

	// get zone id
	for _, v := range respZoneList {
		if reqBody.ZoneName == v.Name {
			zoneID = v.ID
			break
		}
	}

	if zoneID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   "Zone not exist",
		})
	}

	req := &customHttp.HttpRequestHelper{
		Method:        cloudflareCDNAPIURI["zone"]["zoneDetails"]["method"],
		URL:           fmt.Sprintf("%s/%s", cloudflareCDNAPIURI["zone"]["zoneDetails"]["url"], zoneID),
		RequestHeader: cloudflareTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err = json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"CloudflareZoneZoneDetail",
			"CloudflareZoneZoneDetail",
			"CloudflareZoneZoneDetail (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// CloudflareZoneDeleteZone | Cloudflare Zone Delete Zone
// @Summary Cloudflare Zone Delete Zone
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareZoneDeleteZoneRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/delete [post]
func CloudflareZoneDeleteZone(c *gin.Context) {
	// if c.ClientIP() != "127.0.0.1" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"html": "Invalid IP address!",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusBadRequest, gin.H{
		"html": "Invalid IP address!",
	})
	return

	var (
		reqBody CloudflareZoneDeleteZoneRequestBody
		resBody CloudflareCDNResponseBody
		zoneID  string
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   reqBody.AuthKey,
		"X-Auth-Email": reqBody.User,
	}

	// get zone list
	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	respZoneList, err := api.ListZones()

	// get zone id
	for _, v := range respZoneList {
		if reqBody.ZoneName == v.Name {
			zoneID = v.ID
			break
		}
	}

	if zoneID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   "Zone not exist",
		})
	}

	req := &customHttp.HttpRequestHelper{
		Method:        cloudflareCDNAPIURI["zone"]["deleteZone"]["method"],
		URL:           fmt.Sprintf("%s/%s", cloudflareCDNAPIURI["zone"]["deleteZone"]["url"], zoneID),
		RequestHeader: cloudflareTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err = json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"CloudflareZoneZoneDetail",
			"CloudflareZoneZoneDetail",
			"CloudflareZoneZoneDetail (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// CloudflareZonePurgeAllFiles | Cloudflare Zone Purge All Files
// @Summary Cloudflare Zone Purge All Files
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareZonePurgeRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/purge/all/files [post]
func CloudflareZonePurgeAllFiles(c *gin.Context) {
	var (
		cloudflarePurgePath         CloudflareZonePurgeRequestBody
		cloudflareZonePurgePath     string
		cloudflareZonePurgeURLHost  string
		cloudflareZonePurgeZoneName string
	)

	if err := c.ShouldBindJSON(&cloudflarePurgePath); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// zone info
	if !strings.HasPrefix(cloudflarePurgePath.Path, "http://") && !strings.HasPrefix(cloudflarePurgePath.Path, "https://") {
		cloudflareZonePurgePath = "https://" + cloudflarePurgePath.Path
	} else {
		cloudflareZonePurgePath = cloudflarePurgePath.Path
	}

	urlParse, err := url.Parse(cloudflareZonePurgePath)
	if err != nil {
		log.Fatal(err)
	}

	cloudflareZonePurgeURLHost = urlParse.Hostname()
	cloudflareZonePurgeZoneName, _ = publicsuffix.EffectiveTLDPlusOne(cloudflareZonePurgeURLHost)

	/* purge all */

	api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// Fetch the zone ID
	id, err := api.ZoneIDByName(cloudflareZonePurgeZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// Purge zone
	purgeAll := cloudflare.PurgeCacheRequest{
		Everything: true,
	}

	resp, err := api.PurgeCache(zone.ID, purgeAll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	/* end of purge all */

	c.JSON(http.StatusOK, resp)
}

// CloudflareZonePurgeFilesByURL | Cloudflare Zone Purge Files by URL
// @Summary Cloudflare Zone Purge Files by URL
// @Id 1
// @Tags Cloudfalre (Zones)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareZonePurgeRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/zone/purge/files/by/url [post]
func CloudflareZonePurgeFilesByURL(c *gin.Context) {
	var (
		cloudflarePurgePath         CloudflareZonePurgeRequestBody
		cloudflareZonePurgePath     string
		cloudflareZonePurgeURLHost  string
		cloudflareZonePurgeZoneName string
	)

	if err := c.ShouldBindJSON(&cloudflarePurgePath); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// zone info
	if !strings.HasPrefix(cloudflarePurgePath.Path, "http://") && !strings.HasPrefix(cloudflarePurgePath.Path, "https://") {
		cloudflareZonePurgePath = "https://" + cloudflarePurgePath.Path
	} else {
		cloudflareZonePurgePath = cloudflarePurgePath.Path
	}

	urlParse, err := url.Parse(cloudflareZonePurgePath)
	if err != nil {
		log.Fatal(err)
	}

	cloudflareZonePurgeURLHost = urlParse.Hostname()
	cloudflareZonePurgeZoneName, _ = publicsuffix.EffectiveTLDPlusOne(cloudflareZonePurgeURLHost)

	/* purge file */

	api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// fetch the zone ID
	id, err := api.ZoneIDByName(cloudflareZonePurgeZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	purgeFile := cloudflare.PurgeCacheRequest{
		Files: []string{
			cloudflareZonePurgePath,
		},
	}

	purgeResp, err := api.PurgeCache(zone.ID, purgeFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// /* end of purge file */

	c.JSON(http.StatusOK, purgeResp)
}
