package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInsert struct {
	Cl   *mongo.Client
	Db   string
	Coll string
	Data map[string]interface{}
}

type MongoFind struct {
	Cl       *mongo.Client
	Db       string
	Coll     string
	FindData []map[string]interface{}
}

type MongoUpdate struct {
	Cl          *mongo.Client
	Db          string
	Coll        string
	Upsert      bool
	FindData    []map[string]interface{}
	UpdateField string
	UpdateData  interface{}
}

type MongoDelete struct {
	Cl       *mongo.Client
	Db       string
	Coll     string
	FindData []map[string]interface{}
}
