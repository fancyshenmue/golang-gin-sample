package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	customHttp "github.com/fancyshenmue/golang-gin-sample/pkg/http"
	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"github.com/gin-gonic/gin"
)

// CloudflareAccountListAccounts | Cloudflare Account List Accounts
// @Summary Cloudflare Account List Accounts
// @Id 1
// @Tags Cloudfalre (Account)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareAccountListAccountsRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/account/list/accounts [post]
func CloudflareAccountListAccounts(c *gin.Context) {
	var (
		reqBody CloudflareAccountListAccountsRequestBody
		resBody CloudflareListAccountResultTop
	)

	c.ShouldBindJSON(&reqBody)

	var cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   CloudFlareAuthKey,
		"X-Auth-Email": CloudFlareAuthMail,
	}

	req := &customHttp.HttpRequestHelper{
		Method:        cloudflareCDNAPIURI["account"]["listAccounts"]["method"],
		URL:           fmt.Sprintf("%s?page=%s&per_page=%s&direction=%s", cloudflareCDNAPIURI["account"]["listAccounts"]["url"], reqBody.Page, reqBody.PerPage, reqBody.Direction),
		RequestHeader: cloudflareTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"CloudflareAccountListAccounts",
			"CloudflareAccountListAccounts",
			"CloudflareAccountListAccounts (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}

// CloudflareAccountDetails | Cloudflare Account Details
// @Summary Cloudflare Account Details
// @Id 1
// @Tags Cloudfalre (Account)
// @version 1.0
// @accept  json
// @produce json
// @param params body CloudflareAccountDetailRequestBody true "body"
// @Success 200 string string
// @Router /api/v1/cloudflare/account/details [post]
func CloudflareAccountDetails(c *gin.Context) {
	var (
		reqBody CloudflareAccountDetailRequestBody
		resBody CloudflareAccountDetailsTop
	)

	c.ShouldBindJSON(&reqBody)

	var cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   CloudFlareAuthKey,
		"X-Auth-Email": CloudFlareAuthMail,
	}

	req := &customHttp.HttpRequestHelper{
		Method:        cloudflareCDNAPIURI["account"]["details"]["method"],
		URL:           fmt.Sprintf("%s/%s", cloudflareCDNAPIURI["account"]["details"]["url"], reqBody.ID),
		RequestHeader: cloudflareTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"CloudflareAccountDetails",
			"CloudflareAccountDetails",
			"CloudflareAccountDetails (Unmarshal JSON error)",
			err,
		)
	}

	c.JSON(http.StatusOK, resBody)
}
