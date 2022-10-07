package server

import (
	"net/http"

	"github.com/comment-anything/prototype1/templates"
)

type CommentResponse struct {
	*Controller
	Url string
}

// GetCommentsPage serves the Comment Requester page to the user
func (s *Server) GetCommentsPage(w http.ResponseWriter, r *http.Request) {
	controller := r.Context().Value(CtxController).(*Controller)
	comresp := CommentResponse{controller, ""}
	templates.CommentsView.Execute(w, comresp)
}

// PostCommentsPageRequestForURL handles a POST request when a user submits a specific URL they want the comments for.
func (s *Server) PostCommentsPageRequestForURL(w http.ResponseWriter, r *http.Request) {
	// url := r.FormValue("Url")

}
