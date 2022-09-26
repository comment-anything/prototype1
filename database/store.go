package database

import (
	"database/sql"
	"fmt"

	"github.com/comment-anything/prototype1/database/generated"
	"github.com/comment-anything/prototype1/util"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sql.DB
	Queries *generated.Queries
}

// New returns a new Store struct ready to be connected to.
func New() Store {
	var store Store
	return store
}

// NewConnect is a convenience method that calls the Connect() method of a Store before returning it.
func NewConnect() Store {
	store := New()
	store.Connect()
	return store
}

// Connect connects a Store to a Postgres database using the configured settings. It also initializes Store.Queries for database access.
func (s *Store) Connect() {
	postgres, err := sql.Open("postgres", util.Config.DB.ConnectString())
	if err != nil {
		fmt.Println(" Error connecting to postgres : ", err.Error())
	} else {
		s.db = postgres
		s.Queries = generated.New(s.db)
	}
}

// Disconnect disconnects a database from the Store and clears the Queries object.
func (s *Store) Disconnect() {
	s.db.Close()
	s.Queries = nil
	fmt.Println(" Disconnected from Database.")
}
