package postgres

import "os"

var (
	dbUser = os.Getenv("PG_DB_USER")
	dbPass = os.Getenv("PG_DB_PASS")
	dbAddr = os.Getenv("PG_DB_ADDR")
	dbPort = os.Getenv("PG_DB_PORT")
	dbName = os.Getenv("PG_DB_NAME")
)
