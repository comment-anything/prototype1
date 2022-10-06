// The Controller Manager holds a map of initialized controllers and is responsible for initializing and loading controllers
package server

import (
	"context"
	"net/http"
)

type CManager struct {
	Server *Server
	Active map[int64]Controller
}

func NewCManager(server *Server) CManager {
	var cm CManager
	cm.Server = server
	cm.Active = make(map[int64]Controller)
	return cm
}

func (c *CManager) NewController(id int64) (*Controller, error) {
	controller, exists := c.Active[id]
	if !exists {
		user, err := c.Server.DB.Queries.GetUserByUserId(context.TODO(), id)
		if err != nil {
			return nil, err
		}
		controller = Controller{User: user, Manager: c}
		c.Active[id] = controller
	}
	return &controller, nil
}

type contextKey string

const (
	CtxUserid     = contextKey("user-id")
	CtxController = contextKey("controller")
)

func (c *CManager) AttachControllerToContext(next func(wnext http.ResponseWriter, rnext *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(CtxUserid).(int64)
		controller, err := c.NewController(id)
		if err != nil {
			c.Server.GetInvalidPath(w, r)
			return
		}
		newctx := context.WithValue(r.Context(), CtxController, controller)
		next(w, r.WithContext(newctx))
	})

}
