package server

import (
	"net/http"

	"github.com/comment-anything/prototype1/communication"
)

type GuestController struct {
	UserControllerBase
}

func (c *GuestController) HandleCommandBan(msg *communication.Ban, server *Server)                 {}
func (c *GuestController) HandleCommandChangeEmail(msg *communication.ChangeEmail, server *Server) {}
func (c *GuestController) HandleCommandChangeFeedback(msg *communication.ChangeFeedback, server *Server) {
}
func (c *GuestController) HandleCommandChangePassword(msg *communication.SetNewPass, server *Server) {
}
func (c *GuestController) HandleCommandChangeProfileBlurb(msg *communication.ChangeProfileBlurb, server *Server) {
}
func (c *GuestController) HandleCommandCommentReply(msg *communication.CommentReply, server *Server) {
}
func (c *GuestController) HandleCommandCommentVote(msg *communication.CommentVote, server *Server) {}
func (c *GuestController) HandleCommandFeedback(msg *communication.Feedback, server *Server)       {}
func (c *GuestController) HandleCommandGetComments(msg *communication.GetComments, server *Server) {
	c.OnPage = server.PageManager.MoveGuestToPage(c, msg.Url, server)
	if c.OnPage != nil {
		comments := c.OnPage.GetComments(msg.SortedBy, msg.SortAscending)
		packet := communication.CreateCommentsPacket(comments)
		c.nextResponse = append(c.nextResponse, packet)
	}
}
func (c *GuestController) HandleCommandLogout(server *Server) {}
func (c *GuestController) HandleCommandGetUserProfile(msg *communication.GetUserProfile, server *Server) {
}
func (c *GuestController) HandleCommandModerate(msg *communication.Moderate, server *Server) {}
func (c *GuestController) HandleCommandPasswordResetCode(msg *communication.PasswordResetCode, server *Server) {
}
func (c *GuestController) HandleCommandCommentReport(msg *communication.CommentReport, server *Server) {
}
func (c *GuestController) HandleCommandRequestValidation(msg *communication.RequestValidation, server *Server) {
}
func (c *GuestController) HandleCommandValidate(msg *communication.Validate, server *Server) {}
func (c *GuestController) HandleCommandViewBans(msg *communication.ViewBans, server *Server) {}
func (c *GuestController) HandleCommandViewLogs(msg *communication.ViewLogs, server *Server) {}
func (c *GuestController) HandleCommandViewModRecords(msg *communication.ViewModRecords, server *Server) {
}
func (c *GuestController) HandleCommandViewMods(msg *communication.ViewMods, server *Server) {}

func (c *GuestController) Respond(r http.Request, w http.ResponseWriter) {
	c.dispatchResponse(r, w)
}
