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

// CloudflareDNSListDNSRecords | Cloudflare DNS List DNS Records
// @Summary Cloudflare DNS List DNS Records
// @Id 1
// @Tags Cloudfalre (DNS)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareDNSListDNSRecordsRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/dns/list/records [post]
func CloudflareDNSListDNSRecords(c *gin.Context) {
	var (
		reqBody   CloudflareDNSListDNSRecordsRequestBody
		dnsRecord cloudflare.DNSRecord
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get zone id
	zoneID, err := api.ZoneIDByName(reqBody.ZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get zone dns record
	resp, err := api.DNSRecords(zoneID, dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}

// CloudflareDNSCreateRecord | Cloudflare DNS Create Record
// @Summary Cloudflare DNS Create Record
// @Id 1
// @Tags Cloudfalre (DNS)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareDNSCreateDNSRecordRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/dns/create/record [post]
func CloudflareDNSCreateRecord(c *gin.Context) {
	var (
		reqBody CloudflareDNSCreateDNSRecordRequestBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get zone id
	zoneID, err := api.ZoneIDByName(reqBody.ZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// set dns record
	dnsRecord := cloudflare.DNSRecord{
		Type:    reqBody.RecordType,
		Name:    reqBody.RecordName,
		Content: reqBody.RecordContent,
		Proxied: reqBody.Proxied,
	}

	// create dns record
	resp, err := api.CreateDNSRecord(zoneID, dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}

// CloudflareDNSRecordDetail | Cloudflare DNS Record Detail
// @Summary Cloudflare DNS Record Detail
// @Id 1
// @Tags Cloudfalre (DNS)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareDNSRecordDetailRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/dns/record/detail [post]
func CloudflareDNSRecordDetail(c *gin.Context) {
	var (
		reqBody        CloudflareDNSRecordDetailRequestBody
		dnsRecord      cloudflare.DNSRecord
		dnsRecordID    string
		RecordName     string
		RecordHostName string
		ZoneName       string
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get zone name
	if !strings.HasPrefix(reqBody.RecordName, "http://") && !strings.HasPrefix(reqBody.RecordName, "https://") {
		RecordName = "https://" + reqBody.RecordName
	} else {
		RecordName = reqBody.RecordName
	}

	urlParse, err := url.Parse(RecordName)
	if err != nil {
		log.Fatal(err)
	}

	RecordHostName = urlParse.Hostname()
	ZoneName, _ = publicsuffix.EffectiveTLDPlusOne(RecordHostName)

	// get zone id
	zoneID, err := api.ZoneIDByName(ZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get record id
	recordID, err := api.DNSRecords(zoneID, dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	for _, v := range recordID {
		if reqBody.RecordName == v.Name {
			dnsRecordID = v.ID
			break
		}
	}

	resp, err := api.DNSRecord(zoneID, dnsRecordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}

// CloudflareDNSDeleteDNSRecord | Cloudflare DNS Delete DNS Record
// @Summary Cloudflare DNS Delete DNS Record
// @Id 1
// @Tags Cloudfalre (DNS)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareDNSDeleteDNSRecordRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/dns/delete/record [post]
func CloudflareDNSDeleteDNSRecord(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"html": "Invalid IP address!",
	})
	return

	var (
		reqBody        CloudflareDNSDeleteDNSRecordRequestBody
		dnsRecord      cloudflare.DNSRecord
		dnsRecordID    string
		RecordName     string
		RecordHostName string
		ZoneName       string
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	api, err := cloudflare.New(reqBody.AuthKey, reqBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get zone name
	if !strings.HasPrefix(reqBody.RecordName, "http://") && !strings.HasPrefix(reqBody.RecordName, "https://") {
		RecordName = "https://" + reqBody.RecordName
	} else {
		RecordName = reqBody.RecordName
	}

	urlParse, err := url.Parse(RecordName)
	if err != nil {
		log.Fatal(err)
	}

	RecordHostName = urlParse.Hostname()
	ZoneName, _ = publicsuffix.EffectiveTLDPlusOne(RecordHostName)

	// get zone id
	zoneID, err := api.ZoneIDByName(ZoneName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	// get record id
	recordID, err := api.DNSRecords(zoneID, dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	for _, v := range recordID {
		if reqBody.RecordName == v.Name {
			dnsRecordID = v.ID
			break
		}
	}

	err = api.DeleteDNSRecord(zoneID, dnsRecordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info":   err.Error(),
		})
	}

	resp := map[string]string{
		"status":      "success",
		"record ID":   dnsRecordID,
		"record name": reqBody.RecordName,
	}

	c.JSON(http.StatusOK, resp)
}
