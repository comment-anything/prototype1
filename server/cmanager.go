// The Controller Manager holds a map of initialized controllers and is responsible for initializing and loading controllers
package server

import (
	"context"
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

func (c *CManager) Controller(id int64) (*Controller, error) {
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
