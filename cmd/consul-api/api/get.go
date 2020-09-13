package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKey | Get Key
// @Summary Get Key
// @Id 1
// @Tags GET Key
// @version 1.0
// @accept  json
// @produce json
// @param params body ConsulGetKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/getKey [post]
func GetKey(c *gin.Context) {
	var (
		reqBody ConsulGetKeyTop
		resBody CommonInterface
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Get a handle to the KV API
	kv := consulClient.KV()

	// Lookup the pair
	pair, _, err := kv.Get(reqBody.Key, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	json.Unmarshal(pair.Value, &resBody)

	c.JSON(http.StatusOK, resBody)
}

// ListKey | List Key
// @Summary List Key
// @Id 1
// @Tags List Key
// @version 1.0
// @accept  json
// @produce json
// @param params body ConsulAPIListKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/listKey [post]
func ListKey(c *gin.Context) {
	var (
		reqBody      ConsulAPIListKeyTop
		resBodySlice CommonStringSlice
		resBody      CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Get a handle to the KV API
	kv := consulClient.KV()

	// Get Key List
	pair, _, err := kv.List(reqBody.Key, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, v := range pair {
		resBodySlice = append(resBodySlice, v.Key)
	}

	resBody = make(map[string]interface{})
	resBody["keyList"] = resBodySlice

	c.JSON(http.StatusOK, resBody)
}
