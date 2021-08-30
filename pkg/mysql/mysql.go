package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(DBUser, DBPass, DBAddr, DBPort, DBName string) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		""+DBUser+":"+DBPass+"@tcp("+DBAddr+":"+DBPort+")/"+DBName+"?multiStatements=true",
	)
	if err != nil {
		log.Printf("%v\n", err)
	}

	return db, err
}
