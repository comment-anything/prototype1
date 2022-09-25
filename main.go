package main

import (
	"fmt"
	"log"

	"github.com/comment-anything/prototype1/database"
	"github.com/comment-anything/prototype1/server"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("\n Prototype Started. ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure there is a well-formatted .env file in the root of the project. See https://github.com/joho/godotenv and the readme for more information.")
	}
	database.BuildConnectorAndConnect()
	server.StartServer()
	database.Disconnect()
	fmt.Println(" Prototype Ended. ")
}
