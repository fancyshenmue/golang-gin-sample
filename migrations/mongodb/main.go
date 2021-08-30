package main

import (
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	customMongo "golang-gin-sample/pkg/mongo"

	log "github.com/sirupsen/logrus"
)

type Seed struct {
	Collection string `json:"collection"`
	Data       Data   `json:"data"`
}

type Data struct {
	Env         string      `json:"env"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	AppDownload AppDownload `json:"appDownload"`
	Domain      []Domain    `json:"domain"`
}

type AppDownload struct {
	Doamin []string `json:"doamin"`
	Type   string   `json:"type"`
}

type Domain struct {
	Domain []string `json:"domain"`
	Type   string   `json:"type"`
}

var (
	// viper
	V = viper.New()

	// mongodb
	mongoURI    string
	DBMigrate   int
	mongoClient *mongo.Client
	err         error
	mongoDB     string
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	V.SetConfigType("yaml")
	V.SetConfigFile(os.Getenv("_CONFIG_FILE"))
	err := V.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "Config",
		}).Fatalf("Read Config Failed: %v\n", err)
	}

	mongoURI = V.GetStringMapString("mongo")["addr"]
	Credential := options.Credential{
		Username: V.GetStringMapString("mongo")["user"],
		Password: V.GetStringMapString("mongo")["pass"],
	}
	mongoDB = V.GetStringMapString("mongo")["name"]

	mongoClient, err = customMongo.ConnectToMongo(mongoURI, Credential)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MongoDB",
		}).Fatalf("Connect MongoDB Failed: %v\n", err)
	}
}

func main() {
	var seed = []*Seed{
		{
			Collection: "demo",
			Data: Data{
				Env:  "demo",
				Name: "demo",
				Type: "demo",
				AppDownload: AppDownload{
					Doamin: []string{
						"android.demo.com",
						"ios2.demo.com",
					},
					Type: "app",
				},
				Domain: []Domain{
					{
						Domain: []string{
							"www.demo.com",
							"www.demo2.com",
							"www.demo3.com",
						},
						Type: "web",
					},
					{
						Domain: []string{
							"api.demo.com",
							"api.demo2.com",
							"api.demo3.com",
						},
						Type: "api",
					},
				},
			},
		},
		{
			Collection: "demo",
			Data: Data{
				Env:  "demo2",
				Name: "demo2",
				Type: "demo2",
				AppDownload: AppDownload{
					Doamin: []string{
						"android2.demo.com",
						"ios5.demo.com",
					},
					Type: "app",
				},
				Domain: []Domain{
					{
						Domain: []string{
							"www.demo.com",
							"www.demo8.com",
							"www.demo5.com",
						},
						Type: "web",
					},
					{
						Domain: []string{
							"api.demo.com",
							"api.demo8.com",
							"api.demo5.com",
						},
						Type: "api",
					},
				},
			},
		},
	}
	// seed := &Seed{
	// 	Collection: "demo",
	// 	Data: Data{
	// 		Env:  "demo2",
	// 		Name: "demo2",
	// 		AppDownload: AppDownload{
	// 			Doamin: []string{
	// 				"android.demo.com",
	// 				"ios2.demo.com",
	// 			},
	// 			Type: "app",
	// 		},
	// 		Domain: []Domain{
	// 			{
	// 				Domain: []string{
	// 					"www.demo.com",
	// 					"www.demo2.com",
	// 					"www.demo3.com",
	// 				},
	// 				Type: "web",
	// 			},
	// 			{
	// 				Domain: []string{
	// 					"api.demo.com",
	// 					"api.demo2.com",
	// 					"api.demo3.com",
	// 				},
	// 				Type: "api",
	// 			},
	// 		},
	// 	},
	// }

	// findData := make([]map[string]interface{}, 0)
	// findData = append(findData, map[string]interface{}{
	// 	"env": "demo",
	// })
	// findData = append(findData, map[string]interface{}{
	// 	"name": "demo",
	// })

	for _, v := range seed {
		findData := []map[string]interface{}{
			{
				"env": v.Data.Env,
			},
			{
				"name": v.Data.Name,
			},
		}
		updateField := "domain"
		updateData := v.Data

		query := customMongo.MongoUpdate{
			Cl:          mongoClient,
			Db:          mongoDB,
			Coll:        "demo",
			Upsert:      true,
			FindData:    findData,
			UpdateField: updateField,
			UpdateData:  updateData,
		}
		query.UpdateDocument()
	}
}
