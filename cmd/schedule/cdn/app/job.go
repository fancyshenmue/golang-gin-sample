package app

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/cloudflare/cloudflare-go"
	"go.mongodb.org/mongo-driver/bson"

	customHttp "golang-gin-sample/pkg/http"
	customLogHandle "golang-gin-sample/pkg/loghandle"
	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
)

// SyncCloudflareZoneList | Sync Cloudflare Zone list
func SyncCloudflareZoneList() {
	// postgres query
	var (
		// filter
		filtercheck bson.M
		filter      []bson.M

		zoneData bson.M
		zoneList []bson.M
	)

	api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
	if err != nil {
		log.Printf("Initial Cloudflare error: %v", err)
	}

	resp, err := api.ListZones()

	for _, v := range resp {
		zoneData = make(map[string]interface{})
		zoneData["id"] = v.ID
		zoneData["name"] = v.Name
		zoneData["development_mode"] = v.DevMode
		zoneData["original_name_servers"] = v.OriginalNS
		zoneData["original_registrar"] = v.OriginalRegistrar
		zoneData["original_dnshost"] = v.OriginalDNSHost
		zoneData["created_on"] = v.CreatedOn
		zoneData["modified_on"] = v.ModifiedOn
		zoneData["name_servers"] = v.NameServers
		zoneData["owner"] = v.Owner
		zoneData["status"] = v.Status
		zoneData["paused"] = v.Paused
		zoneData["type"] = v.Type
		zoneData["host"] = v.Host
		zoneData["vanity_name_servers"] = v.VanityNS
		zoneData["betas"] = v.Betas
		zoneData["deactivation_reason"] = v.DeactReason
		zoneData["meta"] = v.Meta
		zoneData["account"] = v.Account
		zoneData["verification_key"] = v.VerificationKey
		zoneList = append(zoneList, zoneData)
	}

	for _, v := range zoneList {
		// cleanup
		filter = nil
		filtercheck = make(map[string]interface{})

		// filter
		filtercheck["id"] = fmt.Sprintf("%v", v["id"])
		filtercheck["name"] = fmt.Sprintf("%v", v["name"])
		filter = append(filter, filtercheck)

		// check data exist or not
		query := customMongoDatabase.MongoResultComplexQueryHelper{
			Cl:       mongoClient,
			Db:       customMongoSetup.MongoDatabase,
			Coll:     CloudFlareZoneInfoCollection,
			FindData: filter,
		}

		check, _ := query.FindDataComplexQueryFromMongo()

		// remove old data
		if len(check) >= 1 {
			query.DeleteDataComplexFromMongo()
		}

		// insert new data
		customMongoDatabase.InsertSingleDocument(
			mongoClient,
			customMongoSetup.MongoDatabase,
			CloudFlareZoneInfoCollection,
			v,
		)
	}
}

// SyncCloudflareZoneRecord | Sync Cloudflare Zone Record
func SyncCloudflareZoneRecord() {
	// postgres query
	var (
		// filter
		filtercheck bson.M
		filter      []bson.M

		// Zone Record
		m        bson.M
		zoneList []bson.M

		dnsRecord cloudflare.DNSRecord
	)

	// get cloudflare zone list
	api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
	if err != nil {
		log.Printf("Initial Cloudflare error: %v", err)
	}

	resp, err := api.ListZones()

	for _, zone := range resp {
		m = make(map[string]interface{})
		m["id"] = zone.ID
		zoneList = append(zoneList, m)
	}

	// get zone record
	for _, v := range zoneList {
		api, err := cloudflare.New(CloudFlareAuthKey, CloudFlareAuthMail)
		if err != nil {
			log.Printf("Initial Cloudflare error: %v", err)
		}

		// get zone dns record
		resp, err := api.DNSRecords(fmt.Sprintf("%s", v["id"]), dnsRecord)
		if err != nil {
			log.Printf("Get Cloudflare Zone Record error: %v", err)
		}

		for _, v := range resp {
			// cleanup
			filter = nil
			filtercheck = make(map[string]interface{})
			m = make(map[string]interface{})

			// filter
			filtercheck["type"] = v.Type
			filtercheck["name"] = v.Name
			filtercheck["content"] = v.Content
			filter = append(filter, filtercheck)

			m["id"] = v.ID
			m["type"] = v.Type
			m["name"] = v.Name
			m["content"] = v.Content
			m["proxiable"] = v.Proxiable
			m["proxied"] = v.Proxied
			m["ttl"] = v.TTL
			m["locked"] = v.Locked
			m["zone_id"] = v.ZoneID
			m["zone_name"] = v.ZoneName
			m["created_on"] = v.CreatedOn
			m["modified_on"] = v.ModifiedOn
			m["data"] = v.Data
			m["meta"] = v.Meta
			m["priority"] = v.Priority

			// check data exist or not
			query := customMongoDatabase.MongoResultComplexQueryHelper{
				Cl:       mongoClient,
				Db:       customMongoSetup.MongoDatabase,
				Coll:     CloudFlareZoneRecordInfoCollection,
				FindData: filter,
			}

			check, _ := query.FindDataComplexQueryFromMongo()

			// remove old data
			if len(check) == 1 {
				query.DeleteDataComplexFromMongo()
			}

			// insert new data
			customMongoDatabase.InsertSingleDocument(
				mongoClient,
				customMongoSetup.MongoDatabase,
				CloudFlareZoneRecordInfoCollection,
				m,
			)
		}
	}
}

// SyncVerizonHTTPLargeList | Sync Verizon HTTP Large list
func SyncVerizonHTTPLargeList() {
	var (
		resBody []VerizoneCDNCnameHttpLargeListResponse

		// filter
		filtercheck bson.M
		filter      []bson.M

		// Zone Record
		m             bson.M
		httpLargeList []bson.M
	)

	// get http large info
	req := &customHttp.HttpRequestHelper{
		Method:        verizoneCDNAPIURI["edgeCNAME"]["getAllHTTPLarge"]["method"],
		URL:           fmt.Sprintf("%s", verizoneCDNAPIURI["edgeCNAME"]["getAllHTTPLarge"]["url"]),
		RequestHeader: verizonCDNTokenHeader,
		RequestBody:   nil,
	}

	resp := req.HttpRequest()

	err := json.Unmarshal([]byte(resp), &resBody)
	if err != nil {
		customLogHandle.ErrorHandle(
			"VerizonEdgeCNAMEListHTTPLarge",
			"VerizonEdgeCNAMEListHTTPLarge",
			"VerizonEdgeCNAMEListHTTPLarge (Unmarshal JSON error)",
			err,
		)
	}

	for _, v := range resBody {
		m = make(map[string]interface{})
		m["DirPath"] = v.DirPath
		m["EnableCustomReports"] = v.EnableCustomReports
		m["Id"] = v.ID
		m["MediaTypeId"] = v.MediaTypeID
		m["Name"] = v.Name
		m["OriginId"] = v.OriginID
		m["OriginString"] = v.OriginString
		httpLargeList = append(httpLargeList, m)
	}

	for _, v := range httpLargeList {
		// cleanup
		filter = nil
		filtercheck = make(map[string]interface{})

		// filter
		filtercheck["Name"] = fmt.Sprintf("%v", v["Name"])
		filter = append(filter, filtercheck)

		// check data exist or not
		query := customMongoDatabase.MongoResultComplexQueryHelper{
			Cl:       mongoClient,
			Db:       customMongoSetup.MongoDatabase,
			Coll:     VerizonHTTPLargeInfoCollection,
			FindData: filter,
		}

		check, _ := query.FindDataComplexQueryFromMongo()

		// remove old data
		if len(check) == 1 {
			query.DeleteDataComplexFromMongo()
		}

		// insert new data
		customMongoDatabase.InsertSingleDocument(
			mongoClient,
			customMongoSetup.MongoDatabase,
			VerizonHTTPLargeInfoCollection,
			v,
		)
	}
}
