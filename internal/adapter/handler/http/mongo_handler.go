package http

import (
	"net/http"

	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/usecase/mongo"
	"github.com/gin-gonic/gin"
)

// MongoHandler handles HTTP requests for MongoDB operations
type MongoHandler struct {
	useCase *mongo.MongoUseCase
}

// NewMongoHandler creates a new MongoDB handler
func NewMongoHandler(useCase *mongo.MongoUseCase) *MongoHandler {
	return &MongoHandler{
		useCase: useCase,
	}
}

// GetDocument handles getting documents with filter
// @Summary Get Document
// @Tags MongoDB
// @Accept json
// @Produce json
// @Param body body entity.MongoQuery true "Query parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/getDocument [post]
func (h *MongoHandler) GetDocument(c *gin.Context) {
	var query entity.MongoQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.useCase.GetDocuments(c.Request.Context(), &query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": result,
	})
}

// InsertSingleDocument handles inserting a single document
// @Summary Insert Single Document
// @Tags MongoDB
// @Accept json
// @Produce json
// @Param body body entity.MongoInsert true "Insert parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/insertSingleDocument [post]
func (h *MongoHandler) InsertSingleDocument(c *gin.Context) {
	var insert entity.MongoInsert
	if err := c.ShouldBindJSON(&insert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedID, err := h.useCase.InsertDocument(c.Request.Context(), &insert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"inserted_id": insertedID,
	})
}

// InsertManyDocument handles inserting multiple documents
// @Summary Insert Many Documents
// @Tags MongoDB
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Insert parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/insertManyDocument [post]
func (h *MongoHandler) InsertManyDocument(c *gin.Context) {
	var reqBody struct {
		Database   string        `json:"database" binding:"required"`
		Collection string        `json:"collection" binding:"required"`
		Documents  []interface{} `json:"documents" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedIDs, err := h.useCase.InsertManyDocuments(c.Request.Context(), reqBody.Database, reqBody.Collection, reqBody.Documents)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"inserted_ids": insertedIDs,
	})
}

// UpdateSingleDocument handles updating a single document
// @Summary Update Single Document
// @Tags MongoDB
// @Accept json
// @Produce json
// @Param body body entity.MongoUpdate true "Update parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/updateSingleDocument [put]
func (h *MongoHandler) UpdateSingleDocument(c *gin.Context) {
	var update entity.MongoUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.UpdateDocument(c.Request.Context(), &update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// DeleteSingleDocument handles deleting a single document
// @Summary Delete Single Document
// @Tags MongoDB
// @Accept json
// @Produce json
// @Param body body entity.MongoDelete true "Delete parameters"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/deleteSingleDocument [delete]
func (h *MongoHandler) DeleteSingleDocument(c *gin.Context) {
	var del entity.MongoDelete
	if err := c.ShouldBindJSON(&del); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.useCase.DeleteDocument(c.Request.Context(), &del)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
