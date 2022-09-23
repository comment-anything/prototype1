package views

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	//http.HandleFunc("/user/", getUser)
	//http.HandleFunc("/register/", getRegister)

	server_port := ":" + os.Getenv("SERVER_PORT")
	if server_port == "" {
		fmt.Println(" Environment variable SERVER_PORT must be specified. ")
		panic(" Bad SERVER_PORT env variable. ")
	}
	fmt.Println(" Server starting on port " + server_port)
	log.Fatal(http.ListenAndServe(server_port, nil))
}
