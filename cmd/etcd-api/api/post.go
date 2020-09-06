package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "go.etcd.io/etcd/clientv3"

	customEtcd "golang-gin-sample/pkg/etcd/pkg"
	customEtcdSetup "golang-gin-sample/pkg/etcd/setup"
	customLogHandle "golang-gin-sample/pkg/loghandle"
)

// PutKey | Put Key
// @Summary Put Key
// @Id 1
// @Tags Put Key
// @version 1.0
// @accept  json
// @produce json
// @param params body PutETCdKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/putKey [post]
func (e *Environment) PutKey(c *gin.Context) {
	var (
		ctx = context.TODO()
		// env       = e.Env
		body      map[string]string
		etcdKey   string
		etcdValue string
		m         map[string]interface{}
	)

	c.BindJSON(&body)

	etcdKey = body["key"]
	etcdValue = body["value"]

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo, customEtcdSetup.ETCDConnectUser, customEtcdSetup.ETCDConnectPassword)
	defer cli.Close()

	putresp, err := kv.Put(ctx, etcdKey, etcdValue)
	if err != nil {
		customLogHandle.ErrorHandle(
			"PutKey",
			"Put ETCD Key",
			"Put ETCD Key Error",
			err,
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info":   "Put ETCD Key Error",
		})
	}

	customLogHandle.LogInfo(
		"PutKey",
		"Put ETCD Key",
		fmt.Sprintf("Put ETCD Key %v, Value %v", etcdKey, etcdValue),
	)

	m = make(map[string]interface{})
	m["key"] = etcdKey
	m["value"] = etcdValue
	m["etcdResp"] = putresp

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"info:":  m,
	})
}
