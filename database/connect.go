package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type dbConnector struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var DBConnector dbConnector

// GetConnectionDataFromEnv loads the database connection info from the .env file in the project root and terminates the program if unable to do so.
func GetConnectionDataFromEnv() {
	DBConnector.host = os.Getenv("DB_HOST")
	if DBConnector.host == "" {
		badEnvTerminate("DB_HOST")
	}
	DBConnector.port = os.Getenv("DB_PORT")
	if DBConnector.port == "" {
		badEnvTerminate("DB_PORT")
	}
	DBConnector.user = os.Getenv("DB_USER")
	if DBConnector.user == "" {
		badEnvTerminate("DB_USER")
	}
	DBConnector.password = os.Getenv("DB_PASSWORD")
	if DBConnector.password == "" {
		badEnvTerminate("DB_PASSWORD")
	}
	DBConnector.dbname = os.Getenv("DB_DATABASE_NAME")
	if DBConnector.dbname == "" {
		badEnvTerminate("DB_DATABASE_NAME")
	}
	fmt.Println(" DB env variables loaded.")
}
func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}

// Connect connects to the database and returns the sql database object. After being queried, callers should call `Close` on the returned sql.DB.
func (db dbConnector) Connect() *sql.DB {
	psqlstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.user, db.password, db.dbname)
	postgres, err := sql.Open("postgres", psqlstring)
	if err != nil {
		postgres.Close()
		panic(err)
	}
	return postgres
}
