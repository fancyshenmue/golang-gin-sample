package entity

import "go.mongodb.org/mongo-driver/bson"

// MongoDocument represents a generic MongoDB document
type MongoDocument struct {
	Database   string
	Collection string
	Filter     bson.M
	Data       interface{}
	Limit      int64
}

// MongoQuery represents a query operation on MongoDB
type MongoQuery struct {
	Database   string      `json:"database" binding:"required"`
	Collection string      `json:"collection" binding:"required"`
	Filter     bson.M      `json:"filter"`
	Limit      int64       `json:"limit"`
	Projection interface{} `json:"projection"`
}

// MongoInsert represents an insert operation on MongoDB
type MongoInsert struct {
	Database   string      `json:"database" binding:"required"`
	Collection string      `json:"collection" binding:"required"`
	Document   interface{} `json:"document" binding:"required"`
}

// MongoUpdate represents an update operation on MongoDB
type MongoUpdate struct {
	Database   string      `json:"database" binding:"required"`
	Collection string      `json:"collection" binding:"required"`
	Filter     bson.M      `json:"filter" binding:"required"`
	Update     interface{} `json:"update" binding:"required"`
	Upsert     bool        `json:"upsert"`
}

// MongoDelete represents a delete operation on MongoDB
type MongoDelete struct {
	Database   string `json:"database" binding:"required"`
	Collection string `json:"collection" binding:"required"`
	Filter     bson.M `json:"filter" binding:"required"`
}
