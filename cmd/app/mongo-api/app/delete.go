package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customMongo "golang-gin-sample/pkg/mongo"
)

// DeleteDocument | Delete Document
// @Summary Delete Document
// @Id 4
// @Tags Delete Data
// @version 1.0
// @accept  json
// @produce json
// @param params body DeleteDocumentBody true "body"
// @Success 200 string string
// @Router /api/v1/deleteDocument [delete]
func DeleteDocument(c *gin.Context) {
	var (
		reqBody DeleteDocumentBody
		resBody DeleteDocumentBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check data exist or not
	query := customMongo.MongoDelete{
		Cl:       mongoClient,
		Db:       mongoDB,
		Coll:     reqBody.Collection,
		FindData: reqBody.FindData,
	}

	err := query.DeleteDocument()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Fail",
			"Info":   "Document Not Exist",
		})
	}

	resBody = DeleteDocumentBody{
		Status:     "Success",
		Collection: reqBody.Collection,
		FindData:   reqBody.FindData,
	}

	c.JSON(http.StatusOK, resBody)
}
