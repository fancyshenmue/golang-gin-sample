package http

import (
	"net/http"

	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/usecase/etcd"
	"github.com/gin-gonic/gin"
)

// EtcdHandler handles HTTP requests for etcd operations
type EtcdHandler struct {
	useCase *etcd.EtcdUseCase
}

// NewEtcdHandler creates a new etcd handler
func NewEtcdHandler(useCase *etcd.EtcdUseCase) *EtcdHandler {
	return &EtcdHandler{
		useCase: useCase,
	}
}

// GetKey handles getting a key-value pair
// @Summary Get Key
// @Tags Etcd
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Query parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/getKey [post]
func (h *EtcdHandler) GetKey(c *gin.Context) {
	var query entity.KeyValueQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, err := h.useCase.GetKey(c.Request.Context(), query.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"key":    query.Key,
		"value":  value,
	})
}

// GetKeyWithPrefix handles getting keys with a prefix
// @Summary Get Keys with Prefix
// @Tags Etcd
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Query parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/getKeyWithPrefix [post]
func (h *EtcdHandler) GetKeyWithPrefix(c *gin.Context) {
	var query entity.KeyValueQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	values, err := h.useCase.GetKeysWithPrefix(c.Request.Context(), query.Prefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": values,
	})
}

// PutKey handles storing a key-value pair
// @Summary Put Key
// @Tags Etcd
// @Accept json
// @Produce json
// @Param body body entity.KeyValuePut true "Put parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/putKey [post]
func (h *EtcdHandler) PutKey(c *gin.Context) {
	var put entity.KeyValuePut
	if err := c.ShouldBindJSON(&put); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.PutKey(c.Request.Context(), put.Key, put.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// DeleteKey handles deleting a key
// @Summary Delete Key
// @Tags Etcd
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Delete parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/deleteKey [delete]
func (h *EtcdHandler) DeleteKey(c *gin.Context) {
	var query entity.KeyValueQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.DeleteKey(c.Request.Context(), query.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
