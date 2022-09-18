package views

import (
	"net/http"

	"github.com/comment-anything/prototype1/database"
	"github.com/comment-anything/prototype1/templates"
)

type formError struct {
	ErrorString string
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			templates.RegistrationView.Execute(w, formError{ErrorString: "Failed to parse form."})
			return
		}
		password := r.Form.Get("Password")
		if len(password) < 4 {
			templates.RegistrationView.Execute(w, formError{ErrorString: "Password is too short."})
			return
		}
		password2 := r.Form.Get("Password2")
		if password != password2 {
			templates.RegistrationView.Execute(w, formError{ErrorString: "Passwords must match."})
			return
		}
		username := r.Form.Get("Username")
		if len(username) < 1 {
			templates.RegistrationView.Execute(w, formError{ErrorString: "Username is too short."})
			return
		}
		email := r.Form.Get("Email")
		if len(email) < 1 {
			templates.RegistrationView.Execute(w, formError{ErrorString: "Invalid email."})
			return
		}
		_, err = database.CreateUser(username, email, database.UALPoster, 1, password)
		if err != nil {
			templates.RegistrationView.Execute(w, formError{ErrorString: err.Error()})
			return
		}
		templates.RegistrationView.Execute(w, nil)
	} else {
		templates.RegistrationView.Execute(w, nil)
	}
}
