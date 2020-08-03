package api

import (
	"fmt"
	"os"
)

/* verizon cdn const */
const (
	verizonCustomReportsDisable = 0
	verizonCustomReportsEnable  = 1
	verizonMediaTypeIDHTTPLarge = 3
	verizonMediaTypeIDHTTPSmall = 8
	verizonMediaTypeIDADN       = 14
)

/* end of verizon cdn const */

/* verizon cdn token */

var (
	verizonCDNCommonHeader = map[string]string{
		"Authorization": fmt.Sprintf("TOK:%s", os.Getenv("VERIZON_CDN_TOKEN")),
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Host":          "api.edgecast.com",
	}
	verizonCDNTokenHeader = map[string]string{
		"Authorization": fmt.Sprintf("TOK:%s", os.Getenv("VERIZON_CDN_TOKEN")),
		"Accept":        "application/json",
		"Host":          "api.edgecast.com",
	}
	verizonCDNTokenHeaderDELETE = map[string]string{
		"Authorization": fmt.Sprintf("TOK:%s", os.Getenv("VERIZON_CDN_TOKEN")),
		"Host":          "api.edgecast.com",
	}
)

/* end of verizon cdn token */

/* response body */

type VerizoneCDNResponseBody map[string]interface{}

/* end of response body */

/* verizon cdn api uri map */

var (
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

/* end of verizon cdn api uri map */

/* verizone cdn cname */

type VerizoneCDNCnameHttpLargeListResponse struct {
	DirPath             string `json:"DirPath"`
	EnableCustomReports int64  `json:"EnableCustomReports"`
	ID                  int64  `json:"Id"`
	MediaTypeID         int64  `json:"MediaTypeId"`
	Name                string `json:"Name"`
	OriginID            int64  `json:"OriginId"`
	OriginString        string `json:"OriginString"`
}

type VerizoneCDNCnameAddEdgeCNAMEsRequest struct {
	DirPath             string `json:"DirPath"`
	EnableCustomReports int64  `json:"EnableCustomReports,omitempty"`
	MediaTypeID         int64  `json:"MediaTypeId"`
	Name                string `json:"Name" binding:"required"`
	OriginID            int64  `json:"OriginId"`
}

type VerizoneCDNCnameAddEdgeCNAMEsResponse struct {
	CNameID int64 `json:"CNameId"`
}

/* end of verizone cdn cname */

/* verizon cdn customer origins */

type VerizoneCDNCustomerOriginAddHTTPLargeRequest struct {
	DirectoryName      string                                                     `json:"DirectoryName" binding:"required"`
	HostHeader         string                                                     `json:"HostHeader" binding:"required"`
	HTTPHostnames      []VerizoneCDNCustomerOriginAddHTTPLargeRequestHTTPHostname `json:"HttpHostnames" binding:"required"`
	HTTPLoadBalancing  string                                                     `json:"HttpLoadBalancing" binding:"required"`
	HTTPSHostnames     []VerizoneCDNCustomerOriginAddHTTPLargeRequestHTTPHostname `json:"HttpsHostnames" binding:"required"`
	HTTPSLoadBalancing string                                                     `json:"HttpsLoadBalancing" binding:"required"`
	ShieldPOPS         []VerizoneCDNCustomerOriginAddHTTPLargeRequestShieldPOP    `json:"ShieldPOPs,omitempty"`
}

type VerizoneCDNCustomerOriginAddHTTPLargeRequestHTTPHostname struct {
	Name string `json:"Name"`
}

type VerizoneCDNCustomerOriginAddHTTPLargeRequestShieldPOP struct {
	POPCode string `json:"POPCode,omitempty"`
}

/* end of verizon cdn customer origins */

/* verizon cdn cache */

type VerizoneCDNCachePurgeContentRequest struct {
	MediaPath string `json:"MediaPath" binding:"required"`
	MediaType int64  `json:"MediaType,omitempty"`
}

/* end of verizon cdn cache */
