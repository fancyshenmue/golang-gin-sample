package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnectPostgres | Connect Postgres
func ConnectPostgres() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://"+dbUser+":"+dbPass+"@"+dbAddr+":"+dbPort+"/"+dbName+"?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
