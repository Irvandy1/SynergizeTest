package config

import (
	"database/sql"
	"fmt"
)

func initPostgres() (db *sql.DB, err error) {
	const (
		host     = "localhost"
		port     = 9000
		user     = "postgres"
		password = "password"
		dbname   = "synergizw"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")
	return
}
