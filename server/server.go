/*
Package server uses the net/http library to run a web server which listens for HTTPRequests on ports and serves HTTPResponses. The package is responsible for configuring routing to API end points and setting up the middleware tree to authenticate requests. It calls  package database to interface with a database as needed.

It uses `gorilla/mux` for routing.
*/
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// StartServer configures all routing then starts listening on the configured port.
//
// # Uses
//
// [gorilla/mux]: https://github.com/gorilla/mux for routing.
func StartServer() {

	// Instantiate the gorilla/mux router.
	r := mux.NewRouter()

	// Serve the static index page on root.
	r.Handle("/", http.FileServer(http.Dir("./views")))

	// Serve static assets like images and css.

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	server_port := ":" + os.Getenv("SERVER_PORT")
	if server_port == "" {
		fmt.Println(" Environment variable SERVER_PORT must be specified. ")
		panic(" Bad SERVER_PORT env variable. ")
	}
	fmt.Println(" Server starting on port " + server_port)
	log.Fatal(http.ListenAndServe(server_port, r))
}