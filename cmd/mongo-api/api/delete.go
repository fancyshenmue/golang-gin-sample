package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"
	mongoDatabase "github.com/fancyshenmue/golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "github.com/fancyshenmue/golang-gin-sample/pkg/mongo/setup"
)

// DeleteSingleDocument | Delete Single Document
// @Summary Delete Single Document
// @Id 1
// @Tags Delete Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body MongoDeleteSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/deleteSingleDocument/{collection} [delete]
func DeleteSingleDocument(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var data bson.M

	json.Unmarshal(body, &data)

	// check data exist or not
	query := mongoDatabase.MongoResultHelper{
		Cl:        mongoClient,
		Db:        customMongoSetup.MongoDatabase,
		Coll:      c.Param("collection"),
		FindField: "name",
		FindData:  fmt.Sprintf("%v", data["name"]),
	}

	check, _ := query.FindOneDataFromMongo()
	customLogHandle.ErrorHandle(
		"DeleteSingleDocument",
		"DeleteSingleDocument",
		"DeleteSingleDocument (Get Mongo Document error)",
		err,
	)

	if len(check) == 1 {
		query.DeleteOneDataFromMongo()

		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Fail",
			"info":   "Document Not Exist",
		})
	}
}
