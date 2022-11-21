package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/comment-anything/prototype1/communication"
	"github.com/comment-anything/prototype1/database/generated"
)

type UserControllerInterface interface {
	HandleCommandBan(*communication.Ban, *Server)
	HandleCommandChangeEmail(*communication.ChangeEmail, *Server)
	HandleCommandChangeFeedback(*communication.ChangeFeedback, *Server)
	HandleCommandChangePassword(*communication.SetNewPass, *Server)
	HandleCommandChangeProfileBlurb(*communication.ChangeProfileBlurb, *Server)
	HandleCommandCommentReply(*communication.CommentReply, *Server)
	HandleCommandCommentVote(*communication.CommentVote, *Server)
	HandleCommandFeedback(*communication.Feedback, *Server)
	HandleCommandGetComments(*communication.GetComments, *Server)
	HandleCommandGetUserProfile(*communication.GetUserProfile, *Server)
	HandleCommandLogout(*Server)
	HandleCommandModerate(*communication.Moderate, *Server)
	HandleCommandPasswordResetCode(*communication.PasswordResetCode, *Server)
	HandleCommandCommentReport(*communication.CommentReport, *Server)
	HandleCommandRequestValidation(*communication.RequestVerification, *Server)
	HandleCommandValidate(*communication.Verify, *Server)
	HandleCommandViewBans(*communication.ViewBans, *Server)
	HandleCommandViewDomainReport(*communication.ViewDomainReport, *Server)
	HandleCommandViewUsersReport(*communication.ViewUsersReport, *Server)
	HandleCommandViewLogs(*communication.ViewLogs, *Server)
	HandleCommandViewModRecords(*communication.ViewModRecords, *Server)
	HandleCommandViewMods(*communication.ViewMods, *Server)

	Respond(r http.Request, w http.ResponseWriter)
	GetCurrentPage() *Page
	dispatchResponse(r http.Request, w http.ResponseWriter)
}

type UserControllerBase struct {
	User             generated.User
	lastTokenRefresh time.Time
	//Manager UserManager
	OnPage       *Page
	nextResponse [][]byte
}

// Respond is called by Server on the Controller at the end of the HTTP Response Life Cycle. All built up responses are written to the HTTP Response which is dispatched.
func (u *UserControllerBase) dispatchResponse(r http.Request, w http.ResponseWriter) {
	bytes, err := json.Marshal(u.nextResponse)
	if err != nil {
		w.Write(bytes)
	}
	u.nextResponse = [][]byte{}
}

func (u *UserControllerBase) GetCurrentPage() *Page {
	return u.OnPage
}
