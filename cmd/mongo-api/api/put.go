package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	customLogHandle "golang-gin-sample/pkg/loghandle"
	mongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// UpdateSingleDocument | Update Single Document
// @Summary Update Single Document
// @Id 1
// @Tags Update Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body MongoUpdateSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/updateSingleDocument/{collection} [put]
func UpdateSingleDocument(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var data MongoUpdateSingleDocumentBodyTop

	json.Unmarshal(body, &data)

	// check data exist or not
	query := mongoDatabase.MongoResultHelper{
		Cl:       mongoClient,
		Db:       customMongoSetup.MongoDatabase,
		Coll:     c.Param("collection"),
		FindData: data.Name,
	}

	check, _ := query.FindOneDataFromMongo()
	customLogHandle.ErrorHandle(
		"UpdateSingleDocument",
		"UpdateSingleDocument",
		"UpdateSingleDocument (Get Mongo Document error)",
		err,
	)

	if len(check) != 0 {
		update := mongoDatabase.MongoUpdateHelper{
			Cl:         mongoClient,
			Db:         customMongoSetup.MongoDatabase,
			Coll:       c.Param("collection"),
			FindData:   data.Name,
			Field:      data.Field,
			UpdateData: data.Info,
		}

		update.UpdateSingleDocumentFromMongo()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
	})
}
