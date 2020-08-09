package api

import (
	"fmt"
	"os"
)

/* postgres */

type CloudflarePostgresAuthInfo struct {
	User    string `db:"user" json:"user"`
	AuthKey string `db:"auth_key" json:"auth_key"`
}

/* end of postgres */

/* cloudflare auth info */
var (
	CloudFlareAuthKey  = os.Getenv("CLOUDFLARE_AUTHKEY")
	CloudFlareAuthMail = os.Getenv("CLOUDFLARE_EMAIL")
)

/* end of cloudflare auth info */

/* cloudflare account list */

type CloudflareListAccountResultTop struct {
	Result     []CloudflareListAccountResult   `json:"result"`
	ResultInfo CloudflareListAccountResultInfo `json:"result_info"`
	Success    bool                            `json:"success"`
	Errors     []interface{}                   `json:"errors"`
	Messages   []interface{}                   `json:"messages"`
}

type CloudflareListAccountResult struct {
	ID          string                           `json:"id"`
	Name        string                           `json:"name"`
	Type        string                           `json:"type"`
	Settings    CloudflareListAccountSettings    `json:"settings"`
	LegacyFlags CloudflareListAccountLegacyFlags `json:"legacy_flags"`
	CreatedOn   string                           `json:"created_on"`
}

type CloudflareListAccountLegacyFlags struct {
	EnterpriseZoneQuota CloudflareListAccountEnterpriseZoneQuota `json:"enterprise_zone_quota"`
}

type CloudflareListAccountEnterpriseZoneQuota struct {
	Maximum   int64 `json:"maximum"`
	Current   int64 `json:"current"`
	Available int64 `json:"available"`
}

type CloudflareListAccountSettings struct {
	EnforceTwofactor     bool        `json:"enforce_twofactor"`
	AccessApprovalExpiry interface{} `json:"access_approval_expiry"`
}

type CloudflareListAccountResultInfo struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalPages int64 `json:"total_pages"`
	Count      int64 `json:"count"`
	TotalCount int64 `json:"total_count"`
}

/* end of cloudflare account list */

/* cloudflare account details */

type CloudflareAccountDetailsTop struct {
	Result   CloudflareAccountDetailsResult `json:"result"`
	Success  bool                           `json:"success"`
	Errors   []interface{}                  `json:"errors"`
	Messages []interface{}                  `json:"messages"`
}

type CloudflareAccountDetailsResult struct {
	ID          string                              `json:"id"`
	Name        string                              `json:"name"`
	Type        string                              `json:"type"`
	Settings    CloudflareAccountDetailsSettings    `json:"settings"`
	LegacyFlags CloudflareAccountDetailsLegacyFlags `json:"legacy_flags"`
	CreatedOn   string                              `json:"created_on"`
}

type CloudflareAccountDetailsLegacyFlags struct {
	EnterpriseZoneQuota CloudflareAccountDetailsEnterpriseZoneQuota `json:"enterprise_zone_quota"`
}

type CloudflareAccountDetailsEnterpriseZoneQuota struct {
	Maximum   int64 `json:"maximum"`
	Current   int64 `json:"current"`
	Available int64 `json:"available"`
}

type CloudflareAccountDetailsSettings struct {
	EnforceTwofactor     bool        `json:"enforce_twofactor"`
	AccessApprovalExpiry interface{} `json:"access_approval_expiry"`
}

/* end of cloudflare account details */

/* cloudflare api uri map */

var (
	cloudflareAPIURLRoot = "https://api.cloudflare.com"
	cloudflareCDNAPIURI  = map[string]map[string]map[string]string{
		"account": map[string]map[string]string{
			"listAccounts": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/accounts"),
				"method": "GET",
			},
			"details": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/accounts"),
				"method": "GET",
			},
		},
		"zone": map[string]map[string]string{
			"listZones": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "GET",
			},
			"createZone": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "POST",
			},
			"zoneDetails": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "GET",
			},
			"deleteZone": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "DELETE",
			},
			"purgeAllFiles": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "POST",
			},
			"purgeFilesByURL": map[string]string{
				"url":    fmt.Sprintf("%s%s", cloudflareAPIURLRoot, "/client/v4/zones"),
				"method": "POST",
			},
		},
	}
)

/* end of cloudflare api uri map */

/* cloudflare account */

type CloudflareAccountListAccountsRequestBody struct {
	Page      string `json:"page"`
	PerPage   string `json:"per_page"`
	Direction string `json:"direction"`
}

type CloudflareAccountDetailRequestBody struct {
	ID string `json:"id"`
}

/* end of cloudflare account */

/* cloudflare dns */

type CloudflareDNSListDNSRecordsRequestBody struct {
	User     string `json:"user" binding:"required"`
	AuthKey  string `json:"auth_key" binding:"required"`
	ZoneName string `json:"name" binding:"required"`
}

type CloudflareDNSCreateDNSRecordRequestBody struct {
	User          string `json:"user" binding:"required"`
	AuthKey       string `json:"auth_key" binding:"required"`
	ZoneName      string `json:"zone_name" binding:"required"`
	RecordName    string `json:"record_name" binding:"required"`
	RecordType    string `json:"type" binding:"required"`
	RecordContent string `json:"content" binding:"required"`
	TTL           int    `json:"ttl"`
	Priority      int    `json:"priority,omitempty"`
	Proxied       bool   `json:"proxied" binding:"required"`
}

type CloudflareDNSRecordDetailRequestBody struct {
	User       string `json:"user" binding:"required"`
	AuthKey    string `json:"auth_key" binding:"required"`
	RecordName string `json:"record_name" binding:"required"`
}

type CloudflareDNSDeleteDNSRecordRequestBody struct {
	User       string `json:"user" binding:"required"`
	AuthKey    string `json:"auth_key" binding:"required"`
	RecordName string `json:"record_name" binding:"required"`
}

/* end of cloudflare dns */

/* cloudflare zone */

type CloudflareZoneCreateZoneRequestBody struct {
	User      string `json:"user" binding:"required"`
	AuthKey   string `json:"auth_key" binding:"required"`
	Name      string `json:"name" binding:"required"`
	JumpStart bool   `json:"jump_start"`
	Type      string `json:"type"`
}

type CloudflareZoneCreateZoneAccount struct {
	ID string `json:"id"`
}

type CloudflareZoneZoneDetailRequestBody struct {
	User     string `json:"user" binding:"required"`
	AuthKey  string `json:"auth_key" binding:"required"`
	ZoneName string `json:"name" binding:"required"`
}

type CloudflareZoneDeleteZoneRequestBody struct {
	User     string `json:"user" binding:"required"`
	AuthKey  string `json:"auth_key" binding:"required"`
	ZoneName string `json:"name" binding:"required"`
}

type CloudflareZonePurgeRequestBody struct {
	Path string `json:"path"`
}

/* end of cloudflare zone */
