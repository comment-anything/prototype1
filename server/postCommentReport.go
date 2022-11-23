// src/server/postCommentReport.go

package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/comment-anything/prototype1/communication"
	"github.com/comment-anything/prototype1/database/generated"
)

// API Endpoint for https://commentanywhere.net/newReport
func (server *Server) postCommentReport(request *http.Request, writer http.ResponseWriter) {
	// instantiate a new empty report
	report := communication.PostCommentReport{}
	// attempt to read the body of the comment to the report
	err := json.NewDecoder(request.Body).Decode(&report)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		controller := getControllerInterfaceFromContext(request.Context())
		controller.HandleCommandCommentReport(&report, server)
		controller.Respond(request, writer)
	}

}

// What occurs when a Guest attempts to report a comment.
func (c *GuestController) HandleCommandCommentReport(msg *communication.PostCommentReport, server *Server) {
	// create an error message for transmission to the client
	message := communication.Message{
		Success: false, Text: "You must be logged in to report a comment.",
	}
	// convert that message into a packet for front-end parsing
	bytes, err := communication.CreatePacket(message, communication.ServerMessage)
	if err != nil {
		// append the message to the responses the client is waiting on
		_ = append(c.nextResponse, bytes)
	}
}

// What occurs when a logged-in user attempts to report a comment; a record is inserted into the database.
func (c *MemberController) HandleCommandCommentReport(msg *communication.PostCommentReport, server *Server) {
	// create the comment report in the database
	server.DB.Queries.CreateCommentReport(context.Background(), generated.CreateCommentReportParams{
		ReportingUser: c.User.ID,
		Comment:       msg.CommentId,
		Reason:        sql.NullString{String: msg.Reason},
	})

	// create a response message
	message := communication.Message{
		Success: true, Text: "Comment Report submitted.",
	}
	bytes, err := communication.CreatePacket(message, communication.ServerMessage)
	if err != nil {
		// append the message to the responses the client is waiting on.
		_ = append(c.nextResponse, bytes)
	}
}
