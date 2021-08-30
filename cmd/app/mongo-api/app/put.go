package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	customMongo "golang-gin-sample/pkg/mongo"
)

// UpdateDocument | Update Document
// @Summary Update Document
// @Id 3
// @Tags Update Data
// @version 1.0
// @accept  json
// @produce json
// @param params body UpdateDocumentBody true "body"
// @Success 200 string string
// @Router /api/v1/updateDocument [put]
func UpdateDocument(c *gin.Context) {
	var (
		reqBody UpdateDocumentBody
		resBody UpdateDocumentBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check data exist or not
	query := customMongo.MongoUpdate{
		Cl:          mongoClient,
		Db:          mongoDB,
		Coll:        reqBody.Collection,
		Upsert:      reqBody.Upsert,
		FindData:    reqBody.FindData,
		UpdateField: reqBody.UpdateField,
		UpdateData:  reqBody.UpdateData,
	}
	query.UpdateDocument()

	resBody = UpdateDocumentBody{
		Status:      "Success",
		Collection:  reqBody.Collection,
		FindData:    reqBody.FindData,
		Upsert:      reqBody.Upsert,
		UpdateField: reqBody.UpdateField,
		UpdateData:  reqBody.UpdateData,
	}

	c.JSON(http.StatusOK, resBody)
}
