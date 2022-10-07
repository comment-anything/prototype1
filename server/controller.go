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

func (c *Controller) RefreshAuthCookie(w http.ResponseWriter) error {
	tok, err := GetToken(c.User.ID)
	if err != nil {
		return err
	}
	c.lastTokenRefresh = time.Now()
	auth_cookie := http.Cookie{
		Name:    util.Config.Server.JWTCookieName,
		Value:   tok,
		MaxAge:  0,
		Path:    "/",
		Expires: c.lastTokenRefresh.Add(60 * time.Minute),
	}
	http.SetCookie(w, &auth_cookie)
	return nil
}
