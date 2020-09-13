package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"

	customEtcd "golang-gin-sample/pkg/etcd/pkg"
	customEtcdSetup "golang-gin-sample/pkg/etcd/setup"
	customLogHandle "golang-gin-sample/pkg/loghandle"
)

// GetKey | Get Key
// @Summary Get Key
// @Id 1
// @Tags GET Key
// @version 1.0
// @accept  json
// @produce json
// @param params body GetETCdKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/getKey [post]
func (e *Environment) GetKey(c *gin.Context) {
	var (
		ctx     = context.TODO()
		reqBody GetETCdKeyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo, customEtcdSetup.ETCDConnectUser, customEtcdSetup.ETCDConnectPassword)
	defer cli.Close()

	getresp, err := kv.Get(ctx, reqBody.Key)
	if err != nil {
		customLogHandle.ErrorHandle(
			"GetKey",
			"Get ETCD Key",
			"Get ETCD Key Error",
			err,
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info":   "Get ETCD Key Error",
		})
	}

	// key exist
	if getresp.Count == 1 {
		customLogHandle.LogInfo(
			"GetKey",
			"Get ETCD Key",
			fmt.Sprintf("Get ETCD Key %v", string(getresp.Kvs[0].Key)),
		)

		resBody = make(map[string]interface{})
		resBody["key"] = reqBody.Key
		resBody["value"] = string(getresp.Kvs[0].Value)

		c.JSON(http.StatusOK, resBody)
	}

	// key not exist
	if getresp.Count == 0 {
		customLogHandle.LogInfo(
			"GetKey",
			"Get ETCD Key",
			"ETCD Key not Exist",
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info:":  "ETCD Key not Exist",
		})
	}
}

// GetKeyWithPrefix | Get Get Key With Prefix
// @Summary Get Get Key With Prefix
// @Id 1
// @Tags GET Key
// @version 1.0
// @accept  json
// @produce json
// @param params body GetETCdKeyWithPrefixTop true "body"
// @Success 200 string string
// @Router /api/v1/getKeyWithPrefix [post]
func (e *Environment) GetKeyWithPrefix(c *gin.Context) {
	var (
		ctx        = context.TODO()
		reqBody    GetETCdKeyTop
		resBodyMap CommonMap
		resBody    CommonMapSlice
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo, customEtcdSetup.ETCDConnectUser, customEtcdSetup.ETCDConnectPassword)
	defer cli.Close()

	rangeresp, err := kv.Get(ctx, reqBody.Key, clientv3.WithPrefix())
	if err != nil {
		customLogHandle.ErrorHandle(
			"GetKeyStartsWithPrefix",
			"Get ETCD Starts With Prefix",
			"Get ETCD Starts With Prefix Error",
			err,
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info":   "Get ETCD Key Error",
		})
	}

	// prefix exist
	if rangeresp.Count > 0 {
		customLogHandle.LogInfo(
			"GetKeyStartsWithPrefix",
			"Get ETCD Starts With Prefix",
			fmt.Sprintf("Get ETCD Starts With Prefix: %v", reqBody.Key),
		)

		for _, v := range rangeresp.Kvs {
			resBodyMap = make(map[string]interface{})
			resBodyMap["key"] = string(v.Key)
			resBodyMap["value"] = string(v.Value)
			resBody = append(resBody, resBodyMap)
		}

		c.JSON(http.StatusOK, resBody)
	}

	// prefix not exist
	if rangeresp.Count == 0 {
		customLogHandle.LogInfo(
			"GetKeyStartsWithPrefix",
			"Get ETCD Starts With Prefix",
			"ETCD Prefix Key not Exist",
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info:":  "ETCD Prefix Key not Exist",
		})
	}
}
