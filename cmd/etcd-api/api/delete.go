package api

import (
	"context"
	"net/http"

	customEtcd "github.com/fancyshenmue/golang-gin-sample/pkg/etcd/pkg"
	customEtcdSetup "github.com/fancyshenmue/golang-gin-sample/pkg/etcd/setup"
	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"

	"github.com/gin-gonic/gin"
)

// DeleteKey | Delete Key
// @Summary Delete Key
// @Id 1
// @Tags Delete Key
// @version 1.0
// @accept  json
// @produce json
// @param params body DeleteETCdKeyTop true "body"
// @Success 200 string string
// @Router /api/v1/deleteKey [delete]
func (e *Environment) DeleteKey(c *gin.Context) {
	var (
		ctx = context.TODO()
		// env  = e.Env
		body map[string]string
		key  string
	)

	c.BindJSON(&body)

	key = body["key"]

	kv, cli := customEtcd.ConnectETCD(customEtcdSetup.ETCDConnectInfo)
	defer cli.Close()

	_, err := kv.Delete(ctx, key)
	if err != nil {
		customLogHandle.ErrorHandle(
			"DeleteKey",
			"Delete ETCD Key",
			"Delete ETCD Key Error",
			err,
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"info":   "Delete ETCD Key Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"action":  "Delete Key",
		"keyName": key,
	})
}
