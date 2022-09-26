package main

import (
	"fmt"
	"log"

	"github.com/comment-anything/prototype1/server"
	"github.com/comment-anything/prototype1/util"
)

func main() {
	fmt.Println("\n Prototype Started. ")
	util.Config.Load(".env")
	server, err := server.New()
	if err != nil {
		log.Fatalf("Error initializing server. %s", err.Error())
	} else {
		server.Start()
	}
	fmt.Println(" Prototype Ended. ")
}
