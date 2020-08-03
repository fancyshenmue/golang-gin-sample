package api

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go"
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

	c.ShouldBindJSON(&cloudflarePurgePath)

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

	c.ShouldBindJSON(&cloudflarePurgePath)

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
