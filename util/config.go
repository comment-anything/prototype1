package util

/*
Config holds values parsed from the .env file in project root. It is used across the application to configure connections.
*/

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// dbCredentials are stored in the global Config singleton as Config.db. They hold the connection settings for accessing the Postgres database.
type dbCredentials struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

// serverConfig is stored in the global Config singleton as Config.server. It holds the connection settings for the server.
type serverConfig struct {
	port string
}

type config struct {
	db     dbCredentials
	server serverConfig
}

// Config is a global configuration object singleton holding environment variables and other global data.
var Config config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure there is a well-formatted .env file in the root of the project. See https://github.com/joho/godotenv and the readme for more information.")
	}
	Config.loadDBEnv()
	Config.loadServerEnv()
}

// loadDBEnv loads environment variables into the configuration struct. If it fails to load a variable, it terminates the program process. Correct environment variables are required for the server to run.
func (c *config) loadDBEnv() {
	c.db.host = os.Getenv("DB_HOST")
	if c.db.host == "" {
		badEnvTerminate("DB_HOST")
	}
	c.db.port = os.Getenv("DB_HOST_PORT")
	if c.db.port == "" {
		badEnvTerminate("DB_PORT")
	}
	c.db.user = os.Getenv("DB_USER")
	if c.db.user == "" {
		badEnvTerminate("DB_USER")
	}
	c.db.password = os.Getenv("DB_PASSWORD")
	if c.db.password == "" {
		badEnvTerminate("DB_PASSWORD")
	}
	c.db.dbname = os.Getenv("DB_DATABASE_NAME")
	if c.db.dbname == "" {
		badEnvTerminate("DB_DATABASE_NAME")
	}
}

func (c *config) loadServerEnv() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		badEnvTerminate(port)
	} else {
		Config.server.port = ":" + port

	}
}

func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}
