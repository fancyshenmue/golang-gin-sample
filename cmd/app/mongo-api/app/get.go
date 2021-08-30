package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	customMongo "golang-gin-sample/pkg/mongo"
)

// FindDocument | Find Document
// @Summary Find Document
// @Id 1
// @Tags Get Info
// @version 1.0
// @accept  json
// @produce json
// @param params body FindDocumentBody true "body"
// @Success 200 string string
// @Router /api/v1/FindDocument [post]
func FindDocument(c *gin.Context) {
	var (
		reqBody FindDocumentBody
		resBody FindDocumentBody
	)

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := customMongo.MongoFind{
		Cl:       mongoClient,
		Db:       mongoDB,
		Coll:     reqBody.Collection,
		FindData: reqBody.Data,
	}

	resp, err := query.FindDocument()
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB API",
		}).Fatalf("Find Document Failed: %v\n", err)
	}

	resBody = FindDocumentBody{
		Status:     "Success",
		Collection: reqBody.Collection,
		Data:       resp,
	}

	c.JSON(http.StatusOK, resBody)
}
