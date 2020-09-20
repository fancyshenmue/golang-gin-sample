package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// DeleteSingleDocument | Delete Single Document
// @Summary Delete Single Document
// @Id 1
// @Tags Delete Data
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoDeleteSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/deleteSingleDocument [delete]
func DeleteSingleDocument(c *gin.Context) {
	var (
		reqBody MongoDeleteSingleDocumentBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check data exist or not
	query := customMongoDatabase.MongoResultComplexQueryHelper{
		Cl:       mongoClient,
		Db:       customMongoSetup.MongoDatabase,
		Coll:     reqBody.Collection,
		FindData: reqBody.FindData,
	}

	err := query.DeleteDataComplexFromMongo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Fail",
			"info":   "Document Not Exist",
		})
	}

	resBody = make(map[string]interface{})
	resBody["status"] = "success"
	resBody["query_filter"] = reqBody.FindData

	c.JSON(http.StatusOK, resBody)
}
