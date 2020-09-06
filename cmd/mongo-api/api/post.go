package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	customLogHandle "golang-gin-sample/pkg/loghandle"
	mongoDatabase "golang-gin-sample/pkg/mongo/pkg"
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
// @param params body MongoInsertSingleDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/insertSingleDocument/{collection} [post]
func InsertSingleDocument(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var data bson.M

	json.Unmarshal(body, &data)

	query := mongoDatabase.MongoResultHelper{
		Cl:        mongoClient,
		Db:        customMongoSetup.MongoDatabase,
		Coll:      c.Param("collection"),
		FindField: "name",
		FindData:  fmt.Sprintf("%v", data["name"]),
	}

	// check data exist or not
	check, _ := query.FindOneDataFromMongo()
	customLogHandle.ErrorHandle(
		"InsertSingleDocument",
		"InsertSingleDocument",
		"InsertSingleDocument (Get Mongo Document error)",
		err,
	)

	if len(check) == 0 {
		mongoDatabase.InsertSingleDocument(
			mongoClient,
			customMongoSetup.MongoDatabase,
			c.Param("collection"),
			data,
		)

		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"info":   data,
		})
	} else if len(check) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Fail",
			"info":   "data already exist",
		})
	}
}

// InsertManyDocument | Insert Many Document
// @Summary Insert Many Document
// @Id 1
// @Tags Insert Data
// @version 1.0
// @accept  json
// @produce json
// @param collection path string true "collection"
// @param params body MongoInsertManyDocumentBodyTop true "body"
// @Success 200 string string
// @Router /api/v1/insertManyDocument/{collection} [post]
func InsertManyDocument(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	var (
		data           []bson.M
		insertDataInfo []string
	)

	json.Unmarshal(body, &data)

	for _, v := range data {
		// check data exist or not
		query := mongoDatabase.MongoResultHelper{
			Cl:        mongoClient,
			Db:        customMongoSetup.MongoDatabase,
			Coll:      c.Param("collection"),
			FindField: "name",
			FindData:  fmt.Sprintf("%v", v["name"]),
		}

		check, _ := query.FindOneDataFromMongo()
		customLogHandle.ErrorHandle(
			"InsertManyDocument",
			"InsertManyDocument",
			"InsertManyDocument (Get Mongo Document error)",
			err,
		)

		// insert data if not exist
		if len(check) == 0 {
			mongoDatabase.InsertSingleDocument(
				mongoClient,
				customMongoSetup.MongoDatabase,
				c.Param("collection"),
				v,
			)

			insertDataInfo = append(insertDataInfo, fmt.Sprintf("%v", v["name"]))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"info":   insertDataInfo,
	})
}
