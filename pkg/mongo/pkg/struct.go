package pkg

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoResultHelper struct {
	Cl        *mongo.Client
	Db        string
	Coll      string
	FindField string
	FindData  string
}

type MongoResultComplexQueryHelper struct {
	Cl       *mongo.Client
	Db       string
	Coll     string
	FindData []bson.M
}

type MongoUpdateHelper struct {
	Cl         *mongo.Client
	Db         string
	Coll       string
	FindData   string
	Field      string
	UpdateData interface{}
}

/* auth */

type AuthMemberTop struct {
	Name    string   `json:"name"`
	Memeber []string `json:"memeber"`
}

/* end of auth */
