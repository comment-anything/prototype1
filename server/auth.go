package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/comment-anything/prototype1/util"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const (
	CtxController = contextKey("controller")
)

const (
	TokenIDKey  = "i"
	TokenExpKey = "x"
)

// ReadsAuth is a middleware which causes a Controller to be attached to a Request Context if a valid token is present.
func (s *Server) ReadsAuth(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tcook, err := r.Cookie(util.Config.Server.JWTCookieName)
		if err != nil {
			tstring := tcook.Value
			token, err := jwt.Parse(tstring, keyfunc)
			if err != nil {
				next(w, r)
				return
			}
			if !token.Valid {
				next(w, r)
				return
			}
			claims := token.Claims.(jwt.MapClaims)
			expires := claims[TokenExpKey].(string)
			t, err := time.Parse(time.RFC3339, expires)
			if err != nil {
				next(w, r)
				return
			}
			now := time.Now().Unix()
			expires_at := t.Unix()
			if now > expires_at {
				next(w, r)
				return
			}
			id_string := claims[TokenIDKey].(string)
			raw_id_int, err := strconv.Atoi(id_string)
			if err != nil {
				next(w, r)
				return
			}
			user_id := int64(raw_id_int)

			controller, err := s.Controllers.Controller(user_id)
			if err != nil {
				next(w, r)
				return
			}
			controller.RefreshAuthCookie(w)
			newctx := context.WithValue(r.Context(), CtxController, controller)
			next(w, r.WithContext(newctx))
			return
		}
		next(w, r)
	})
}

// keyfunc is used by JWT, it confirms the signing method and returns the secret key for parsing
func keyfunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodECDSA)
	if !ok {
		return nil, errors.New("couldn't parse token: bad signing method")
	}
	return util.Config.Server.JWTKey, nil
}

// GetToken simply returns a JWT token signed with the secret key, with an expiry time of 1 hour and for the userid given as a parameter. It performs no validation.
func GetToken(userid int64) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims[TokenExpKey] = time.Now().Add(60 * time.Minute).Format(time.RFC3339)
	claims[TokenIDKey] = fmt.Sprint(userid)
	tstring, err := token.SignedString(util.Config.Server.JWTKey)
	if err != nil {
		return "", err
	}
	return tstring, nil
}
