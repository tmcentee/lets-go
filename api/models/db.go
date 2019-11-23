package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //idk this calms down golint
)

var db *sql.DB

//InitDB initializes the database singleton and verfies connection
func InitDB(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}

	//If we can't reach the db, panic
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
