package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

// PutKey | Put Key
// @Summary Put Key
// @Id 1
// @Tags Put Key
// @version 1.0
// @accept  json
// @produce json
// @param params body PutConsulKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/putKey [post]
func PutKey(c *gin.Context) {
	var (
		reqBody        PutConsulKeyTop
		valueDataCheck CommonInterface
		resBody        CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check value type
	reqValue := reqBody.Value

	switch valueType := reqValue.(type) {
	case nil:
		fmt.Println("check nil")
		log.Printf("Type is %v", valueType)
		c.JSON(http.StatusBadRequest, "Empty Value")
		return
	case int:
		valueDataCheck = fmt.Sprintf("%s", reqBody.Value)
	case string:
		valueDataCheck = fmt.Sprintf("%s", reqBody.Value)
	case map[string]interface{}:
		valueDataCheck, err = json.Marshal(reqBody.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Marshal JSON Error")
			return
		}
	default:
		c.JSON(http.StatusBadRequest, "Unknown Value Type")
		return
	}

	// Get a handle to the KV API
	kv := consulClient.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: reqBody.Key, Value: []byte(fmt.Sprintf("%s", valueDataCheck))}
	_, err = kv.Put(p, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resBody = make(map[string]interface{})
	resBody["key"] = reqBody.Key
	resBody["value"] = reqBody.Value

	c.JSON(http.StatusOK, resBody)
}
