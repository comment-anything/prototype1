package server

/*

JWT lets us store session information client side, in a sense. Their tokens and claims are parsed against a secret key we keep.


Here is some info about JWT  https://jwt.io/introduction


Currently this file only contains some example functions from a tutorial.


My current thought is that being "authed" should cause a pointer to be attached to the Request Context which gives the users some kind of associated Controller. Middleware and endpoints could make use of that Controller to get and send user-specific data.



*/

import (
	"net/http"
	"time"

	"context"

	"github.com/golang-jwt/jwt"
)

func J() {

	jwt.DecodeSegment("this")
}

// generateJWT is Just an example method for examining how to create a Json Web Token
func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(60 * time.Minute)
	claims["auth_level"] = 1
	claims["user"] = "username" // <-- link to Users... cache or DB?
	//										  secret key would be here not "abc" (env variable)
	tokenString, err := token.SignedString("abc")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// addRefToContext is an example of using middleware to add a pointer to the context of a request.
func addRefToContext(endpointfunc func(w http.ResponseWriter, req *http.Request), userHandle int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rctx := req.Context()
		rctx = context.WithValue(rctx, "UserHandle", userHandle)
		endpointfunc(w, req.WithContext(rctx))
	})
}

// This will wrap the parameterized endpoint function with the JWT verifying middleware
func verifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header["Token"] != nil {
			token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("Can't validate against the wrong cryptography type!!"))
					if err != nil {
						return nil, err
					}
				}
				//     > the secret key goes here again
				return "abc", nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, err2 := w.Write([]byte("Didnt parse against secret key!"))
				if err2 != nil {
					return
				}
			}
			if token.Valid {
				endpointHandler(w, req)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Invalid token!"))
				if err != nil {
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("NO! No token!"))
			if err != nil {
				return
			}
		}
	})
}
