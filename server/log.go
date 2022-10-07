package server

// Logging all requests is a work in progress.
// It needs to wrap a base level handler in the mux router and I'm not yet sure how to do that.

import (
	"log"
	"net/http"

	"github.com/comment-anything/prototype1/util"
)

func ConsoleLogRequests(handler http.Handler) http.Handler {
	if util.Config.Server.DoesLogAll {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received a %s request with %v cookies for %v.\n", r.Method, len(r.Cookies()), r.URL)
			handler.ServeHTTP(w, r)
		})
	}
	return handler
}
