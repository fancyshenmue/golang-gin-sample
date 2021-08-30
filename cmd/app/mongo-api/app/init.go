package app

import (
	customMongo "golang-gin-sample/pkg/mongo"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	V.AddRemoteProvider("consul", os.Getenv("_CONSUL_ENDPOINT"), os.Getenv("_CONSUL_KEY_ENDPOINT"))
	V.SetConfigType(os.Getenv("_CONSUL_CONFIG_TYPE"))
	if err := V.ReadRemoteConfig(); err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "Consul",
		}).Fatalf("Connect Consul Failed: %v\n", err)
		return
	}

	mongoURI := V.GetStringMapString("mongo")["address"]
	Credential := options.Credential{
		Username: V.GetStringMapString("mongo")["user"],
		Password: V.GetStringMapString("mongo")["pass"],
	}
	mongoDB = V.GetStringMapString("mongo")["database"]

	mongoClient, err = customMongo.ConnectToMongo(mongoURI, Credential)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Fatalf("Connect MongoDB Failed: %v\n", err)
	}

	// app
	os.Setenv("APP_PORT", V.GetStringMapString("app")["port"])
}
