package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customLogHandle "golang-gin-sample/pkg/loghandle"
	mongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// GetDocument | Get Document
// @Summary Get Document
// @Id 1
// @Tags Get Info
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoGetBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/getDocument [post]
func GetDocument(c *gin.Context) {
	var (
		body MongoGetBodyTop
	)

	c.BindJSON(&body)

	query := mongoDatabase.MongoResultHelper{
		Cl:        mongoClient,
		Db:        customMongoSetup.MongoDatabase,
		Coll:      body.Collection,
		FindField: body.FindField,
		FindData:  body.FindData,
	}

	resp, err := query.FindOneDataFromMongo()
	customLogHandle.ErrorHandle(
		"GetDocument",
		"GetDocument",
		"GetDocument (Get Mongo Document error)",
		err,
	)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"info":   resp,
	})
}

// GetDocumentRegex | Get Document Regex
// @Summary Get Document Regex
// @Id 1
// @Tags Get Info
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoGetBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/getDocumentRegex [post]
func GetDocumentRegex(c *gin.Context) {
	var (
		body MongoGetBodyTop
	)

	c.BindJSON(&body)

	query := mongoDatabase.MongoResultHelper{
		Cl:        mongoClient,
		Db:        customMongoSetup.MongoDatabase,
		Coll:      body.Collection,
		FindField: body.FindField,
		FindData:  body.FindData,
	}

	resp := query.FindDataRegexFromMongo()
	customLogHandle.ErrorHandle(
		"GetDocumentRegex",
		"GetDocumentRegex",
		"GetDocumentRegex (Get Mongo Document error)",
		err,
	)

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"info":   resp,
	})
}
