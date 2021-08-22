package app

import (
	"fmt"
	customLogHandle "golang-gin-sample/pkg/loghandle"
	customMongo "golang-gin-sample/pkg/mongo/pkg"
	customPostgresSetup "golang-gin-sample/pkg/postgres"

	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "log"
)

func init() {
	V.SetConfigType("yaml")
	V.SetConfigFile(os.Getenv("_CONFIG_FILE"))
	err := V.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	/* service info */
	// cloudflare
	CloudFlareAuthKey = V.GetStringMapString("cloudflare")["cloudflare_authkey"]
	CloudFlareAuthMail = V.GetStringMapString("cloudflare")["cloudflare_email"]
	cloudflareTokenHeader = map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Key":   CloudFlareAuthKey,
		"X-Auth-Email": CloudFlareAuthMail,
	}

	// verizon
	verizonCDNTokenHeader = map[string]string{
		"Authorization": fmt.Sprintf("TOK:%s", V.GetStringMapString("verizon")["verizon_cdn_token"]),
		"Accept":        "application/json",
		"Host":          "api.edgecast.com",
	}
	verizoneAccountNumber = V.GetStringMapString("verizon")["verizon_cdn_account_number"]
	verizonAPIURLRoot = V.GetStringMapString("verizon")["verizon_api_url_root"]
	verizoneCDNAPIURI = map[string]map[string]map[string]string{
		"edgeCNAME": map[string]map[string]string{
			"getAllHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames/httplarge"),
				"method": "GET",
			},
			"getEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "GET",
			},
			"addEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "POST",
			},
			"deleteEdgeCNAME": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "cnames"),
				"method": "DELETE",
			},
		},
		"customerOrigins": map[string]map[string]string{
			"getHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins/httplarge"),
				"method": "GET",
			},
			"addHTTPLarge": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins/httplarge"),
				"method": "POST",
			},
			"delete": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "origins"),
				"method": "DELETE",
			},
		},
		"cache": map[string]map[string]string{
			"purgeContent": map[string]string{
				"url":    fmt.Sprintf("%s/%s/%s", verizonAPIURLRoot, verizoneAccountNumber, "edge/purge"),
				"method": "PUT",
			},
		},
	}

	// schedule
	ScheduleSyncCloudflareZoneListCron = V.GetStringMapString("schedule")["sync_cloudflare_zone_list_cron"]
	ScheduleSyncCloudflareZoneRecord = V.GetStringMapString("schedule")["sync_cloudflare_zone_record"]
	ScheduleSyncVerizonHttpLargeList = V.GetStringMapString("schedule")["sync_verizon_http_large_list"]
	/* end of service info */

	/* log setup */

	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)

	/* end of log setup */

	/* connect mongo */

	MongoURI = fmt.Sprintf("mongodb://%s", V.GetStringMapString("mongo")["mongo_address"])
	MongoCredential := options.Credential{
		Username: V.GetStringMapString("mongo")["mongo_user"],
		Password: V.GetStringMapString("mongo")["mongo_pass"],
	}

	mongoClient, err = customMongo.ConnectToMongo(MongoURI, MongoCredential)
	customLogHandle.ErrorHandle(
		"init",
		"init",
		"init (connect mongo error)",
		err,
	)

	/* end of connect mongo */

	/* connect postgres */

	os.Setenv("PG_DB_USER", V.GetStringMapString("postgres")["pg_db_user"])
	os.Setenv("PG_DB_PASS", V.GetStringMapString("postgres")["pg_db_pass"])
	os.Setenv("PG_DB_ADDR", V.GetStringMapString("postgres")["pg_db_addr"])
	os.Setenv("PG_DB_PORT", V.GetStringMapString("postgres")["pg_db_port"])
	os.Setenv("PG_DB_NAME", V.GetStringMapString("postgres")["pg_db_name"])

	postgresClient, err = customPostgresSetup.ConnectPostgres()
	customLogHandle.ErrorHandle(
		"init",
		"init",
		"init (connect postgres error)",
		err,
	)

	/* end of connect postgres */
}
