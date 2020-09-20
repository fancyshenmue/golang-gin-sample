package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// InsertSingleDocument | Insert Single Document
// @Summary Insert Single Document
// @Id 1
// @Tags Insert Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body MongoAPIInsertSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/insertSingleDocument [post]
func InsertSingleDocument(c *gin.Context) {
	var (
		reqBody MongoAPIInsertSingleDocumentBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := customMongoDatabase.MongoInsertHelper{
		Cl:   mongoClient,
		Db:   customMongoSetup.MongoDatabase,
		Coll: reqBody.Collection,
		Data: reqBody.Body,
	}

	query.InsertSingleDocument()

	resBody = make(map[string]interface{})
	resBody["status"] = "Success"
	resBody["data"] = reqBody.Body

	c.JSON(http.StatusOK, resBody)
}

// InsertManyDocument | Insert Many Document
// @Summary Insert Many Document
// @Id 1
// @Tags Insert Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body MongoAPIInsertManyDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/insertManyDocument [post]
func InsertManyDocument(c *gin.Context) {
	var (
		reqBody MongoAPIInsertManyDocumentBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for _, data := range reqBody.Body {
		query := customMongoDatabase.MongoInsertHelper{
			Cl:   mongoClient,
			Db:   customMongoSetup.MongoDatabase,
			Coll: reqBody.Collection,
			Data: data,
		}

		query.InsertSingleDocument()
	}

	resBody = make(map[string]interface{})
	resBody["status"] = "Success"

	c.JSON(http.StatusOK, resBody)
}
