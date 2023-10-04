package databasePostgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateConn() (*sql.DB, error) {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "pass1234"
	dbname := "pass1234"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)
}
