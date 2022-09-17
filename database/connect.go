package database

import (
	_ "database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConfirmPackage() {
}

type dbConnectionInfo struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var condat dbConnectionInfo

func GetConnectionFromEnv() {
	condat.host = os.Getenv("DB_HOST")
	if condat.host == "" {
		badEnvTerminate("DB_HOST")
	}
	condat.port = os.Getenv("DB_PORT")
	if condat.port == "" {
		badEnvTerminate("DB_PORT")
	}
	condat.user = os.Getenv("DB_USER")
	if condat.user == "" {
		badEnvTerminate("DB_USER")
	}
	condat.password = os.Getenv("DB_PASSWORD")
	if condat.password == "" {
		badEnvTerminate("DB_PASSWORD")
	}
	condat.dbname = os.Getenv("DB_DATABASE_NAME")
	if condat.dbname == "" {
		badEnvTerminate("DB_DATABASE_NAME")
	}
	fmt.Println(" DB env variables loaded.")
}

func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}
