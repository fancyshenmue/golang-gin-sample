package postgres

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnectPostgres | Connect Postgres
func ConnectPostgres() (*sqlx.DB, error) {
	var (
		dbUser = os.Getenv("PG_DB_USER")
		dbPass = os.Getenv("PG_DB_PASS")
		dbAddr = os.Getenv("PG_DB_ADDR")
		dbPort = os.Getenv("PG_DB_PORT")
		dbName = os.Getenv("PG_DB_NAME")
	)

	db, err := sqlx.Connect(
		"postgres",
		"postgres://"+dbUser+":"+dbPass+"@"+dbAddr+":"+dbPort+"/"+dbName+"?sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
