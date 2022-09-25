package server

import (
	"log"
	"net/http"
)

func LogAllRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received a %s request for %v.", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
