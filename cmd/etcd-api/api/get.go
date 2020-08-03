package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"

	customEtcd "github.com/fancyshenmue/golang-gin-sample/pkg/etcd/pkg"
	customEtcdSetup "github.com/fancyshenmue/golang-gin-sample/pkg/etcd/setup"
	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"
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
		ctx = context.TODO()
		// env  = e.Env
		body map[string]string
		key  string
		m    map[string]interface{}
	)

	c.BindJSON(&body)

	key = body["key"]

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo)
	defer cli.Close()

	getresp, err := kv.Get(ctx, key)
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

		m = make(map[string]interface{})
		m["key"] = key
		m["value"] = string(getresp.Kvs[0].Value)

		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"info:":  m,
		})
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
		// env       = e.Env
		body      map[string]string
		keyPreifx string
		m         map[string]interface{}
		resp      []map[string]interface{}
	)

	c.BindJSON(&body)

	keyPreifx = body["prefix"]

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo)
	defer cli.Close()

	rangeresp, err := kv.Get(context.TODO(), keyPreifx, clientv3.WithPrefix())
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
			fmt.Sprintf("Get ETCD Starts With Prefix: %v", keyPreifx),
		)

		for _, v := range rangeresp.Kvs {
			m = make(map[string]interface{})
			m["key"] = string(v.Key)
			m["value"] = string(v.Value)
			resp = append(resp, m)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"info:":  resp,
		})
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
