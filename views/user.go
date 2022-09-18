package views

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/comment-anything/prototype1/database"
	"github.com/comment-anything/prototype1/templates"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	id_s := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(id_s)
	if err == nil {
		user, err := database.GetUser(id)
		if err != nil {
			w.Write([]byte("That user doesn't exist!"))
			return
		}
		_ = templates.UserViewSingle.Execute(w, user)

	} else {
		fmt.Println(err.Error())
	}
}
