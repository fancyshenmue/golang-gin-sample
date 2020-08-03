package api

import (
	"fmt"
	"os"
)

/* grey cdn token */

var (
	greyCDNTokenHeader = map[string]string{
		"Content-Type":  "application/json",
		"greycdn-token": os.Getenv("GREYCDN_TOKEN"),
	}
)

/* end of grey cdn token */

/* grey cdn api uri map */

var (
	greyCDNAPIURLRoot = "https://api2.greypanel.com"
	greyCDNAPIURI     = map[string]map[string]map[string]string{
		"site": map[string]map[string]string{
			"listAll": map[string]string{
				"url":    fmt.Sprintf("%s/%s", greyCDNAPIURLRoot, "/api/v1/site/list/all"),
				"method": "POST",
			},
			"getSiteInfo": map[string]string{
				"url":    fmt.Sprintf("%s/%s", greyCDNAPIURLRoot, "/api/v1/site/view"),
				"method": "GET",
			},
		},
		"cache": map[string]map[string]string{
			"purgeBySite": map[string]string{
				"url":    fmt.Sprintf("%s/%s", greyCDNAPIURLRoot, "/api/v1/cache/purge/by-site"),
				"method": "POST",
			},
			"purgeBySiteURI": map[string]string{
				"url":    fmt.Sprintf("%s/%s", greyCDNAPIURLRoot, "/api/v1/cache/purge/by-site-uri"),
				"method": "POST",
			},
		},
	}
)

/* end of grey cdn api uri map */

/* grey cdn http body */

type GreyCDNSiteListResponseTop struct {
	Result   []GreyCDNSiteListResponse `json:"result"`
	Response string                    `json:"response"`
	Message  string                    `json:"message"`
}

type GreyCDNSiteListResponse struct {
	UID          string `json:"uid"`
	UniqueName   string `json:"uniqueName"`
	Name         string `json:"name"`
	SourceIPS    string `json:"sourceIps"`
	Protocol     string `json:"protocol"`
	Status       int64  `json:"status"`
	ConfigStatus int64  `json:"configStatus"`
	RecordCount  int64  `json:"recordCount"`
	DomainCount  int64  `json:"domainCount"`
	Type         string `json:"type"`
	OpsFlag      int64  `json:"opsFlag"`
	Description  string `json:"description"`
}

// GreyCDNCacheBySiteRequestBody | Grey CDN Cache By Site Request Body
type GreyCDNCacheBySiteRequestBody struct {
	UID string `json:"uid"`
}

// GreyCDNCacheBySiteResponseBody | Grey CDN Cache By Site Response Body
type GreyCDNCacheBySiteResponseBody struct {
	Result   GreyCDNCacheBySiteResponseBodyResult `json:"result"`
	Response string                               `json:"response"`
	Message  string                               `json:"message"`
}

// GreyCDNCacheBySiteResponseBodyResult | Grey CDN Cache By Site Response Body Result
type GreyCDNCacheBySiteResponseBodyResult struct{}

// GreyCDNCacheBySiteURIRequestBody | Grey CDN Cache By Site URI Request Body
type GreyCDNCacheBySiteURIRequestBody struct {
	UID string `json:"uid"`
	URI string `json:"uri"`
}

// GreyCDNCacheBySiteURIResponseBody | Grey CDN Cache By Site URI Response Body
type GreyCDNCacheBySiteURIResponseBody struct {
	Result   GreyCDNCacheBySiteURIResponseBodyResult `json:"result"`
	Response string                                  `json:"response"`
	Message  string                                  `json:"message"`
}

// GreyCDNCacheBySiteURIResponseBodyResult | Grey CDN Cache By Site URI Response Body Result
type GreyCDNCacheBySiteURIResponseBodyResult struct{}

/* end of grey cdn http body */
