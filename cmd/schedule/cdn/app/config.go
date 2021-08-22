package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

/* viper */

var (
	V = viper.New()
)

/* end of viper */

/* cloudflare auth info */

var (
	CloudFlareAuthKey  string
	CloudFlareAuthMail string

	cloudflareTokenHeader map[string]string
)

/* end of cloudflare auth info */

var (
	/* verizon auth info */
	verizonCDNTokenHeader map[string]string
	/* end of verizon auth info */

	verizoneAccountNumber string
	verizonAPIURLRoot     string
	verizoneCDNAPIURI     map[string]map[string]map[string]string
)

/* common */

var (
	err error
)

/* end of common */

/* mongo */

var (
	mongoClient *mongo.Client
	MongoURI    string
)

type CloudflareAuthInfo struct {
	User    string `db:"user" json:"user" binding:"required"`
	AuthKey string `db:"auth_key" json:"auth_key" binding:"required"`
}

/* end of mongo */

/* postgres */

var (
	postgresClient *sqlx.DB
)

/* end of postgres */

/* schedule */

var (
	ScheduleSyncCloudflareZoneListCron string
	ScheduleSyncCloudflareZoneRecord   string
	ScheduleSyncVerizonHttpLargeList   string
)

/* end of schedule */

type VerizoneCDNCnameHttpLargeListResponse struct {
	DirPath             string `json:"DirPath"`
	EnableCustomReports int64  `json:"EnableCustomReports"`
	ID                  int64  `json:"Id"`
	MediaTypeID         int64  `json:"MediaTypeId"`
	Name                string `json:"Name"`
	OriginID            int64  `json:"OriginId"`
	OriginString        string `json:"OriginString"`
}
