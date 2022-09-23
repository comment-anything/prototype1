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

type dbConnector struct {
	Credentials dbCredentials
	Postgres    *sql.DB
}

/** DB will be used throughout the program to connect to the database **/
var DB dbConnector

// BuildConnectorAndConnect loads the database connection info from the .env file in the project root and terminates the program if unable to do so. It then attempts to connect to the database.
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
func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}

func Disconnect() {
	fmt.Println(" Disconnected from Database.")
	DB.Postgres.Close()
}
