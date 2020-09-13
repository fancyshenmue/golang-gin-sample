package api

import "go.mongodb.org/mongo-driver/bson"

/* auth */

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

/* end of auth */

/* api */

type CommonInterface interface{}
type CommonMap map[string]interface{}
type CommonMapSlice []CommonMap

type MongoAPIGetBodyTop struct {
	Collection string   `json:"collection" binding:"required"`
	Body       []bson.M `json:"body" binding:"required"`
}

type MongoAPIGetRegexBodyTop struct {
	Collection string `json:"collection" binding:"required"`
	Field      string `json:"field" binding:"required"`
	Data       string `json:"data" binding:"required"`
}

// MongoAPIInsertSingleDocumentBodyTop
type MongoAPIInsertSingleDocumentBodyTop struct {
	Collection string                 `json:"collection" binding:"required"`
	Body       map[string]interface{} `json:"body" binding:"required"`
}

type MongoAPIInsertManyDocumentBodyTop struct {
	Collection string                   `json:"collection" binding:"required"`
	Body       []map[string]interface{} `json:"body" binding:"required"`
}

type MongoUpdateSingleDocumentBodyTop struct {
	Name  string
	Field string
	Info  interface{}
}

type MongoDeleteSingleDocumentBodyTop struct {
	Name string
}

/* end of api */
