package api

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	customLogHandle "github.com/fancyshenmue/golang-gin-sample/pkg/loghandle"
	customMongo "github.com/fancyshenmue/golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "github.com/fancyshenmue/golang-gin-sample/pkg/mongo/setup"
)

var (
	mongoClient *mongo.Client
	err         error
)

func init() {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)

	mongoClient, err = customMongo.ConnectToMongo(customMongoSetup.MongoURI, customMongoSetup.Credential)
	customLogHandle.ErrorHandle(
		"init",
		"init",
		"init (connect mongo error)",
		err,
	)
}
