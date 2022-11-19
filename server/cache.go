package server

import "github.com/comment-anything/prototype1/communication"

// Page contains data retrieved from the database.
type Page struct {
	// The full path for this page, with both the domain and subsequent path.
	fullPath string
	// An array of CachedComments containing comment data for comments posted on this page.
	comments     map[int64]CachedComment
	usersOnPage  map[int64]UserControllerInterface
	guestsOnPage map[int64]UserControllerInterface
}

func (p *Page) GetComments(sortedBy string, ascending bool) []communication.Comment {
	// todo....
	return []communication.Comment{}
}

// CachedComment contains data on a comment.
type CachedComment struct {
	// The ID of the Cached Comment
	id int64
	// The content of the comment
	content string
	// The ID of the user who posted the comment.
	userId int64
	// The ID of the parent comment.
	parent int64
	// The name of the user who posted the comment.
	username string
	// An array of CachedVote data, representing all votes that have been made on the comment.
	votes []CachedVote
	// The time when this comment was posted.
	createdAt int64
	// Whether the comment is set to hidden or not.
	hidden bool
}

// UpVote takes a userId representing the user voting, a valid category, e.g. "funny", and a value, e.g., -1, 0, 1
func (*CachedComment) Vote(userId int64, category string, value int8) {}

// CachedVote contains cached data about a comment vote.
type CachedVote struct {
	// The ID of the user who cast the vote.
	userId int64
	// The name of the user who cast the vote.
	username string
	// The vote category, e.g., “funny”, “factual”, “agree”
	category string
	// How the user voted; -1 if it was a downvote and 1 if it was an upvote.
	value int8
}
