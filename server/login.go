package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/comment-anything/prototype1/templates"
)

// GetLogin is called when an HTTP GET request is received at the /login endpoint. It serves the login page from a template.
func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	templates.LoginView.Execute(w, "")
}

type LoginError struct {
	ErrorStrings []string
}

func loginErr(ers ...string) LoginError {
	var le LoginError
	le.ErrorStrings = ers
	return le
}

// PostLogin is called when an HTTP POST request is received at the /login endpoint. If the form data fails a registration check, the user is returned to the login page. The template is populated with an error string accordingly.
func (s *Server) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")

	// TODO if there is client-side password encrpytion, decrypt it here
	ctx := context.TODO()
	user, err := s.DB.Queries.GetUserByUserName(ctx, username)
	if err != nil {
		templates.LoginView.Execute(w, loginErr(err.Error()))
		return
	}

	// TODO server-encrypt the password here, before comparison
	if password != user.Password {
		templates.LoginView.Execute(w, loginErr("Invalid password"))
		return
	}

	controller, err := s.Controllers.Controller(user.ID)
	if err != nil {
		GetErrPg(w, r, err.Error(), "Couldn't generate controller")
		return
	}

	err = controller.RefreshAuthCookie(w)
	if err != nil {
		fmt.Println(" ~ login post failed to refresh auth cookie; " + err.Error())
	}
	http.Redirect(w, r, "/authed/dash", http.StatusFound)

	//templates.DashboardView.Execute(w, controller)

	// currently serving dashboard, but probably should redirect to auth route to get user the right URL in header, have auth validate the token that was just generated

}
