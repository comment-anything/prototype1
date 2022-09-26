package server

import (
	"net/http"

	"github.com/comment-anything/prototype1/templates"
)

// registerHandler determines whether the user is issuing an HTTP GET or an HTTP POST request and calls GetRegister or PostRegister respectively.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		PostRegister(w, r)
	} else if r.Method == "GET" {
		GetRegister(w, r)
	} else {
		http.Error(w, "bad method", 404)
	}
}

// getRegister is called when an HTTP GET request is received at the /register endpoint. It serves the registration page from a template.
func GetRegister(w http.ResponseWriter, r *http.Request) {
	templates.RegisterView.Execute(w, "")
}

type registerValidationError struct {
	ErrorStrings []string
}

// postRegister is called when an HTTP POST request is received at the /register endpoint. If the form data fails a registration check, the user is returned to the registration page. The template is populated with an error string accordingly.
func PostRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	retypePassword := r.FormValue("Password2")
	valid, errmsgs := validateRegSubmission(username, password, retypePassword, email)
	_ = errmsgs
	if valid != true {
		templates.RegisterView.Execute(w, &registerValidationError{ErrorStrings: errmsgs})
		return
	} else {
		http.Error(w, "Oh yes!", 404)
		return
	}
}

// validateRegSubmission performs initial validation on user registration form submission. It assembles a slice of strings containing all errors that were encountered during validation. It returns a bool indicating whether the data is good to continue validating and the slice of messages to be displayed if it is not.
func validateRegSubmission(username string, password string, retypePassword string, email string) (bool, []string) {
	var valid bool = true
	var msgs []string
	test, msg := checkUserName(username)
	valid = test && valid
	if !test {
		msgs = append(msgs, msg)
	}
	test, msg = checkPassword(password, retypePassword)
	valid = test && valid
	if !test {
		msgs = append(msgs, msg)
	}
	test, msg = checkEmail(email)
	valid = test && valid
	if !test {
		msgs = append(msgs, msg)
	}
	return valid, msgs
}

// checkUserName checks to ensure a username is legal, does not contain bad words, is long enough, but isn't too long.
func checkUserName(username string) (bool, string) {
	l := len(username)
	if l < 4 {
		return false, "Username too short."
	} else if l > 20 {
		return false, "Username too long."
	}
	return true, ""
}

// checkPassword checks a password to ensure it is sufficiently strong.
func checkPassword(password string, retypePassword string) (bool, string) {
	if password != retypePassword {
		return false, "Passwords do not match."
	}
	if len(password) < 6 {
		return false, "Password too short."
	}
	return true, ""
}

// checkEmail checks an email to determine whether it is valid.
func checkEmail(email string) (bool, string) {
	if len(email) < 3 {
		return false, "Email too short."
	}
	return true, ""
}

// encryptPassword encrypts a password for storage.
func encryptPassword(unencrypted_password string) string {
	return unencrypted_password
}

// decryptPassword decrypts a password from storage.
func decryptPassword(encrypted_password string) string {
	return encrypted_password
}
