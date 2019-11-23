package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //idk this calms down golint
)

//Datastore defines availabe database operations
type Datastore interface {
	AllPlayers() ([]*Player, error)
	AllTeams() ([]*Team, error)
	SinglePlayer(int) (*Player, error)
	SingleTeam(int) (*Team, error)
}

//DB struct wrapper around a sql connection pool
type DB struct {
	*sql.DB
}

//NewDB initializes a db connection and returns a DB struct
func NewDB(connectionString string) (*DB, error) {
	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}

	//If we can't reach the db, panic
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return &DB{db}, nil
}
