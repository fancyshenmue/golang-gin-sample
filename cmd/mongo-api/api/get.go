package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customLogHandle "golang-gin-sample/pkg/loghandle"
	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// GetDocument | Get Document
// @Summary Get Document
// @Id 1
// @Tags Get Info
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoAPIGetBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/getDocument [post]
func GetDocument(c *gin.Context) {
	var (
		reqBody MongoAPIGetBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := customMongoDatabase.MongoResultComplexQueryHelper{
		Cl:       mongoClient,
		Db:       customMongoSetup.MongoDatabase,
		Coll:     reqBody.Collection,
		FindData: reqBody.Body,
	}

	resp, err := query.FindDataComplexQueryFromMongo()
	customLogHandle.ErrorHandle(
		"GetDocument",
		"GetDocument",
		"GetDocument (Get Mongo Document error)",
		err,
	)

	resBody = make(map[string]interface{})
	resBody["status"] = "Success"
	resBody["result"] = resp

	c.JSON(http.StatusOK, resBody)
}

// GetDocumentRegex | Get Document Regex
// @Summary Get Document Regex
// @Id 1
// @Tags Get Info
// @version 1.0
// @accept  json
// @produce json
// @param params body MongoAPIGetRegexBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/getDocumentRegex [post]
func GetDocumentRegex(c *gin.Context) {
	var (
		reqBody MongoAPIGetRegexBodyTop
		resBody CommonMap
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := customMongoDatabase.MongoResultHelper{
		Cl:        mongoClient,
		Db:        customMongoSetup.MongoDatabase,
		Coll:      reqBody.Collection,
		FindField: reqBody.Field,
		FindData:  reqBody.Data,
	}

	resp := query.FindDataRegexFromMongo()
	customLogHandle.ErrorHandle(
		"GetDocumentRegex",
		"GetDocumentRegex",
		"GetDocumentRegex (Get Mongo Document error)",
		err,
	)

	resBody = make(map[string]interface{})
	resBody["status"] = "Success"
	resBody["result"] = resp

	c.JSON(http.StatusOK, resBody)
}
