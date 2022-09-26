package server

// Logging all requests is a work in progress.
// It needs to wrap a base level handler in the mux router and I'm not yet sure how to do that.

import (
	"log"
	"net/http"
)

type f func(w http.ResponseWriter, r *http.Request)

func LogAllRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received a %s request for %v.", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func Log2(wrapped f) f {
	return func(w http.ResponseWriter, r *http.Request) {
		wrapped(w, r)
	}
}
