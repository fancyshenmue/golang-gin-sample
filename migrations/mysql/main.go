package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/viper"
)

var (
	V      = viper.New()
	DbUser string
	DbPass string
	DbAddr string
	DbPort string
	DbName string
)

func init() {
	V.SetConfigType("yaml")
	V.SetConfigFile(os.Getenv("_CONFIG_FILE"))
	err := V.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	DbUser = V.GetStringMapString("mysql")["mysql_db_user"]
	DbPass = V.GetStringMapString("mysql")["mysql_db_pass"]
	DbAddr = V.GetStringMapString("mysql")["mysql_db_addr"]
	DbPort = V.GetStringMapString("mysql")["mysql_db_port"]
	DbName = V.GetStringMapString("mysql")["mysql_db_name"]
}

func main() {
	db, err := sql.Open(
		"mysql",
		""+DbUser+":"+DbPass+"@tcp("+DbAddr+":"+DbPort+")/"+DbName+"?multiStatements=true",
	)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"demo",
		driver,
	)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	m.Steps(2)
}
