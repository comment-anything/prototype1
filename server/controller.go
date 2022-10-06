// A Controller pointer is attached to authed requests. It is used to provide user-specific information to a requester.
package server

import "github.com/comment-anything/prototype1/database/generated"

type Controller struct {
	User    generated.User
	Manager *CManager
}
