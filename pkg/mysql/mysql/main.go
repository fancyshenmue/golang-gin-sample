package main

import (
	"database/sql"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/viper"

	customMySQL "sre-app/pkg/mysql"
)

var (
	V         = viper.New()
	DBUser    string
	DBPass    string
	DBAddr    string
	DBPort    string
	DBName    string
	DBMigrate int

	db *sql.DB
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

	DBUser = V.GetStringMapString("mysql")["user"]
	DBPass = V.GetStringMapString("mysql")["pass"]
	DBAddr = V.GetStringMapString("mysql")["addr"]
	DBPort = V.GetStringMapString("mysql")["port"]
	DBName = V.GetStringMapString("mysql")["name"]
	DBMigrate, _ = strconv.Atoi(V.GetStringMapString("mysql")["migrate_version"])

	db, err = customMySQL.Connect(
		DBUser,
		DBPass,
		DBAddr,
		DBPort,
		DBName,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"status": false,
			"type":   "MySQL",
		}).Fatalf("Connect MySQL Failed: %v\n", err)
	}
}

func main() {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		DBName,
		driver,
	)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	m.Steps(DBMigrate)
}
