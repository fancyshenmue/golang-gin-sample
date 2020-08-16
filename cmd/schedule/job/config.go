package job

import (
	"fmt"
	"os"
)

/* cloudflare auth info */

var (
	CloudFlareAuthKey  = os.Getenv("CLOUDFLARE_AUTHKEY")
	CloudFlareAuthMail = os.Getenv("CLOUDFLARE_EMAIL")

	cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   CloudFlareAuthKey,
		"X-Auth-Email": CloudFlareAuthMail,
	}
)

/* end of cloudflare auth info */

var (
	/* verizon auth info */

	verizonCDNTokenHeader = map[string]string{
		"Authorization": fmt.Sprintf("TOK:%s", os.Getenv("VERIZON_CDN_TOKEN")),
		"Accept":        "application/json",
		"Host":          "api.edgecast.com",
	}

	/* end of verizon auth info */

	verizoneAccountNumber = os.Getenv("VERIZON_CDN_ACCOUNT_NUMBER")
	verizonAPIURLRoot     = "https://api.edgecast.com/v2/mcc/customers"
	verizoneCDNAPIURI     = map[string]map[string]map[string]string{
		"edgeCNAME": map[string]map[string]string{
			"getAllHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames/httplarge"),
				"method": "GET",
			},
			"getEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "GET",
			},
			"addEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "POST",
			},
			"deleteEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "DELETE",
			},
		},
		"customerOrigins": map[string]map[string]string{
			"getHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins/httplarge"),
				"method": "GET",
			},
			"addHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins/httplarge"),
				"method": "POST",
			},
			"delete": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins"),
				"method": "DELETE",
			},
		},
		"cache": map[string]map[string]string{
			"purgeContent": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "edge/purge"),
				"method": "PUT",
			},
		},
	}
)

type VerizoneCDNCnameHttpLargeListResponse struct {
	DirPath             string `json:"DirPath"`
	EnableCustomReports int64  `json:"EnableCustomReports"`
	ID                  int64  `json:"Id"`
	MediaTypeID         int64  `json:"MediaTypeId"`
	Name                string `json:"Name"`
	OriginID            int64  `json:"OriginId"`
	OriginString        string `json:"OriginString"`
}
