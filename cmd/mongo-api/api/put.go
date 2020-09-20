package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// UpdateSingleDocument | Update Single Document
// @Summary Update Single Document
// @Id 1
// @Tags Update Data
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoUpdateSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/updateSingleDocument [put]
func UpdateSingleDocument(c *gin.Context) {
	var (
		reqBody MongoUpdateSingleDocumentBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check data exist or not
	query := customMongoDatabase.MongoUpdateHelper{
		Cl:          mongoClient,
		Db:          customMongoSetup.MongoDatabase,
		Coll:        reqBody.Collection,
		Upsert:      reqBody.Upsert,
		FindData:    reqBody.FindData,
		UpdateField: reqBody.UpdateField,
		UpdateData:  reqBody.UpdateData,
	}

	query.UpdateSingleDocumentFromMongo()

	resBody = make(map[string]interface{})
	resBody["status"] = "Success"
	resBody["collection"] = reqBody.Collection
	resBody["query_filter"] = reqBody.FindData
	resBody["upsert"] = reqBody.Upsert
	resBody["update_field"] = reqBody.UpdateField
	resBody["update_data"] = reqBody.UpdateData

	c.JSON(http.StatusOK, resBody)
}
