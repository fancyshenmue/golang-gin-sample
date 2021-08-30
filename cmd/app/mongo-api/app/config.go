package app

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// consul
	V = viper.New()

	// mongo
	mongoClient *mongo.Client
	err         error
	mongoDB     string
)

type commonMap map[string]interface{}
