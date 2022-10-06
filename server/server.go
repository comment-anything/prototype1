package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/comment-anything/prototype1/database"
	"github.com/comment-anything/prototype1/templates"
	"github.com/comment-anything/prototype1/util"
	"github.com/gorilla/mux"
)

// Server holds the server state including routes and provides methods for server operations.
type Server struct {
	router *mux.Router
	DB     database.Store
}

// New returns a new Server instance with routing applied.
func New() (*Server, error) {
	server := &Server{}
	server.setupRouter()
	server.DB = database.New()
	return server, nil
}

// setupRouter configures the routing for the server.
func (s *Server) setupRouter() {
	// Instantiate the gorilla/mux router.
	r := mux.NewRouter()

	// Handle erroneous requests.
	r.NotFoundHandler = http.HandlerFunc(s.GetInvalidPath)
	r.MethodNotAllowedHandler = http.HandlerFunc(s.GetInvalidPath)

	// Serve static assets from /static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Serve index page on root.
	r.HandleFunc("/", s.GetIndex)
	// Serve register page.
	r.HandleFunc("/register", s.RegisterHandler)

	r.HandleFunc("/login", s.GetLogin).Methods(http.MethodGet)
	r.HandleFunc("/login", s.PostLogin).Methods(http.MethodPost)

	// TODO: Wrap some component of the mux router with the logging function.
	s.router = r
}

// Start causes the server to begin listening on the configured port.
func (s *Server) Start() {
	s.DB.Connect()
	fmt.Println(" Database connection initialized.")
	port := util.Config.Server.Port
	fmt.Println(" Server listening on port", port)
	log.Println(http.ListenAndServe(port, s.router))
	fmt.Println(" Server closed.")
}

// GetIndex serves the home page in response to an http Request.
func (s *Server) GetIndex(w http.ResponseWriter, r *http.Request) {
	templates.IndexView.Execute(w, "")
}

func (s *Server) GetInvalidPath(w http.ResponseWriter, r *http.Request) {
	templates.ErrorView.Execute(w, "")
}
