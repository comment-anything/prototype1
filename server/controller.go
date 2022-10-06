// A Controller pointer is attached to authed requests. It is used to provide user-specific information to a requester.
package server

import (
	"net/http"
	"time"

	"github.com/comment-anything/prototype1/database/generated"
	"github.com/comment-anything/prototype1/util"
)

type Controller struct {
	User             generated.User
	lastTokenRefresh time.Time
	Manager          *CManager
}

func (c *Controller) RefreshAuthCookie(w http.ResponseWriter) {
	tok, _ := GetToken(c.User.ID)
	c.lastTokenRefresh = time.Now()
	auth_cookie := http.Cookie{
		Name:  util.Config.Server.JWTCookieName,
		Value: tok,
	}
	http.SetCookie(w, &auth_cookie)
}
