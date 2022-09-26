package util

/*
Config holds values parsed from the .env file in project root. It is used across the application to configure connections.
*/

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// dbCredentials are stored in the global Config singleton as Config.DB. It holds the connection settings for accessing the Postgres database.
type dbCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

// serverConfig is stored in the global Config singleton as Config.server. It holds the connection settings for the server.
type serverConfig struct {
	Port       string
	DoesLogAll bool
}

type config struct {
	DB     dbCredentials
	Server serverConfig
}

// Config is a global configuration object singleton holding environment variables and other global data.
var Config config

// Load loads the environment variables from the .env file. It should be called in the main function and then in the TestMain function of every package that needs access to those environment variables. While main calls the function with a path to the current working directory, tests will have to use relative directories to find the .env file.
func (c *config) Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file. Ensure there is a well-formatted .env file in the root of the project. See https://github.com/joho/godotenv and the readme for more information.")
	}
	Config.loadDBEnv()
	Config.loadServerEnv()
}

// loadDBEnv loads environment variables into the configuration struct. If it fails to load a variable, it terminates the program process. Correct environment variables are required for the server to run.
func (c *config) loadDBEnv() {
	c.DB.Host = os.Getenv("DB_HOST")
	if c.DB.Host == "" {
		badEnvTerminate("DB_HOST")
	}
	c.DB.Port = os.Getenv("DB_HOST_PORT")
	if c.DB.Port == "" {
		badEnvTerminate("DB_PORT")
	}
	c.DB.User = os.Getenv("DB_USER")
	if c.DB.User == "" {
		badEnvTerminate("DB_USER")
	}
	c.DB.Password = os.Getenv("DB_PASSWORD")
	if c.DB.Password == "" {
		badEnvTerminate("DB_PASSWORD")
	}
	c.DB.DBname = os.Getenv("DB_DATABASE_NAME")
	if c.DB.DBname == "" {
		badEnvTerminate("DB_DATABASE_NAME")
	}
}

// DBString builds a string from the database connection credentials and returns it. For use with sql.Open.
func (d *dbCredentials) ConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.DBname)
}

func (c *config) loadServerEnv() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		badEnvTerminate("SERVER_PORT")
	} else {
		Config.Server.Port = ":" + port
	}
	log := os.Getenv("SERVER_LOG_ALL")
	if log == "" {
		badEnvTerminate("SERVER_LOG_ALL")
	} else if log == "1" || log == "true" || log == "True" {
		c.Server.DoesLogAll = true
	} else {
		c.Server.DoesLogAll = false
	}
}

func badEnvTerminate(name string) {
	log.Fatalf(" Error parsing environment variable %v. Terminating.\n", name)
}
