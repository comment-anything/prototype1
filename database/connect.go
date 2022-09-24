package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type dbCredentials struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

/*
dbConnector holds a reference to the credentials and an instance of a Postgres db.

@Todo: Replace with the generated DBX from /sqlc/db.go?
*/
type dbConnector struct {
	Credentials dbCredentials
	Postgres    *sql.DB
}

/*
* DB will be used throughout the program to connect to the database.

@Todo: Or will it? Replace with DBX from /sqlc/db.go?
*/
var DB dbConnector

// BuildConnectorAndConnect loads the database connection credentials from the .env file in the project root and terminates the program if unable to do so. It then initializes a *sql.DB and attempts to connect to it, storing that database handle in the dbConnector singleton if it succeeds.
func BuildConnectorAndConnect() {
	DB.Credentials.host = os.Getenv("DB_HOST")
	if DB.Credentials.host == "" {
		badEnvTerminate("DB_HOST")
	}
	DB.Credentials.port = os.Getenv("DB_HOST_PORT")
	if DB.Credentials.port == "" {
		badEnvTerminate("DB_PORT")
	}
	DB.Credentials.user = os.Getenv("DB_USER")
	if DB.Credentials.user == "" {
		badEnvTerminate("DB_USER")
	}
	DB.Credentials.password = os.Getenv("DB_PASSWORD")
	if DB.Credentials.password == "" {
		badEnvTerminate("DB_PASSWORD")
	}
	DB.Credentials.dbname = os.Getenv("DB_DATABASE_NAME")
	if DB.Credentials.dbname == "" {
		badEnvTerminate("DB_DATABASE_NAME")
	}
	fmt.Println(" DB env variables loaded.")

	psqlstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB.Credentials.host, DB.Credentials.port, DB.Credentials.user, DB.Credentials.password, DB.Credentials.dbname)
	postgres, err := sql.Open("postgres", psqlstring)
	if err != nil {
		fmt.Println(" Error connecting to postgres. Check your credentials.")
	} else {
		DB.Postgres = postgres
		fmt.Println(" Database Connection established.")
	}
}

// badEnvTerminate is called if an environment variable is not set properly; the program cannot function if it can't connect to the database.
func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}

// Disconnect closes the Postgres database connection.
func Disconnect() {
	fmt.Println(" Disconnected from Database.")
	DB.Postgres.Close()
}
