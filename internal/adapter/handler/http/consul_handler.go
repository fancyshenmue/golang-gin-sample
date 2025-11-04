package http

import (
	"net/http"

	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/usecase/consul"
	"github.com/gin-gonic/gin"
)

// ConsulHandler handles HTTP requests for Consul operations
type ConsulHandler struct {
	useCase *consul.ConsulUseCase
}

// NewConsulHandler creates a new Consul handler
func NewConsulHandler(useCase *consul.ConsulUseCase) *ConsulHandler {
	return &ConsulHandler{
		useCase: useCase,
	}
}

// GetKey handles getting a key-value pair
// @Summary Get Key
// @Tags Consul
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Query parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/getKey [post]
func (h *ConsulHandler) GetKey(c *gin.Context) {
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

// ListKeys handles listing keys with a prefix
// @Summary List Keys
// @Tags Consul
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Query parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/listKey [post]
func (h *ConsulHandler) ListKeys(c *gin.Context) {
	var query entity.KeyValueQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	values, err := h.useCase.ListKeys(c.Request.Context(), query.Prefix)
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
// @Tags Consul
// @Accept json
// @Produce json
// @Param body body entity.KeyValuePut true "Put parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/putKey [post]
func (h *ConsulHandler) PutKey(c *gin.Context) {
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
// @Tags Consul
// @Accept json
// @Produce json
// @Param body body entity.KeyValueQuery true "Delete parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/deleteKey [delete]
func (h *ConsulHandler) DeleteKey(c *gin.Context) {
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
