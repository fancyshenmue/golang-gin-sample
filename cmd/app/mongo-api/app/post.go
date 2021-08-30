package app

import (
	customMongo "golang-gin-sample/pkg/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InsertDocument | Insert Document
// @Summary Insert Document
// @Id 2
// @Tags Insert Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body InsertDocumentBody true "body"
// @Success 200 string string
// @Router /api/v1/insertDocument [post]
func InsertDocument(c *gin.Context) {
	var (
		reqBody InsertDocumentBody
		resBody InsertDocumentBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := customMongo.MongoInsert{
		Cl:   mongoClient,
		Db:   mongoDB,
		Coll: reqBody.Collection,
		Data: reqBody.Data,
	}

	query.InsertDocument()

	resBody = InsertDocumentBody{
		Status:     "Success",
		Collection: reqBody.Collection,
		Data:       reqBody.Data,
	}

	c.JSON(http.StatusOK, resBody)
}

// func InsertDocument() {
// 	x := "The Polyglot Developer Podcast"
// 	data := bson.D{
// 		{Key: "title", Value: x},
// 		{Key: "author", Value: "Nic Raboy"},
// 		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
// 	}

// 	query := customMongo.MongoInsert{
// 		Cl:   mongoClient,
// 		Db:   mongoDB,
// 		Coll: "demo",
// 		Data: data,
// 	}

// 	query.InsertDocument()
// }
