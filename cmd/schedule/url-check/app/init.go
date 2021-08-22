package app

import (
	"log"
	"os"
)

func init() {
	V.AddRemoteProvider("consul", os.Getenv("_CONSUL_ENDPOINT"), os.Getenv("_CONSUL_KEY_ENDPOINT"))
	V.SetConfigType(os.Getenv("_CONSUL_CONFIG_TYPE"))
	if err := V.ReadRemoteConfig(); err != nil {
		log.Println(err)
		return
	}

	// service start
	log.Printf("%s\n", "// Service Check URL Status")

	// service required info
	URLList = V.GetStringSlice("url_list")
	SlackTokenID = V.GetStringMapString("slack")["slack_token_id"]
	SlackChannelID = V.GetStringMapString("slack")["slack_channel_id"]
}
