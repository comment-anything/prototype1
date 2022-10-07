package server

import (
	"net/http"
	"time"

	"github.com/comment-anything/prototype1/util"
)

func (s *Server) GetLogout(w http.ResponseWriter, r *http.Request) {
	delete_cookie := &http.Cookie{
		Name:    util.Config.Server.JWTCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, delete_cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}
