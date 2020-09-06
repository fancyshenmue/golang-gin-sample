package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteKey | Delete Key
// @Summary Delete Key
// @Id 1
// @Tags Delete Key
// @version 1.0
// @accept  json
// @produce json
// @param params body DeleteConsulKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/deleteKey [delete]
func DeleteKey(c *gin.Context) {
	var (
		reqBody DeleteConsulKeyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Get a handle to the KV API
	kv := consulClient.KV()

	_, err = kv.Delete(reqBody.Key, nil)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resBody = make(map[string]interface{})
	resBody["status"] = "Delete Success"
	resBody["key"] = reqBody.Key

	c.JSON(http.StatusOK, resBody)
}
