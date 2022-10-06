package server

import (
	"net/http"

	"github.com/comment-anything/prototype1/templates"
)

// GetDash serves the Dashboard
func (s *Server) GetDash(w http.ResponseWriter, r *http.Request) {
	controller := r.Context().Value(CtxController).(*Controller)
	if controller == nil {
		GetErrPg(w, r, "no controller reference in context")
		return
	}
	templates.DashboardView.Execute(w, controller)
}
