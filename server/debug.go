// for serving an error page with more detailed messages

package server

import (
	"net/http"

	"github.com/comment-anything/prototype1/templates"
)

type dberr struct {
	ErrorMessages []string
}

func dber(strings []string) dberr {
	return dberr{ErrorMessages: strings}
}

func GetErrPg(w http.ResponseWriter, r *http.Request, strings ...string) {
	//serve template
	templates.DebugView.Execute(w, dber(strings))
}
