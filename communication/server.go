package comm

// AdminAccessLog contains data needed by Admins to see an access log.
type AdminAccessLog struct {
	// The IP of anyone accessing the logs.
	ip string
	// The ID of the current log.
	logId int64
	// The full URL that was accessed.
	url string
	// The ID of the user this access record is associated with, if the user was logged in. Otherwise, this field will be empty or 0.
	userId string
	// The username of the accessor If applicable
	username string
}

// AdminDomainReport contains data needed by Admins to see information about activity on a particular domain.
type AdminDomainReport struct {
	// number of comments on a domain.
	commentCount int32
	// the string of the domain.
	domain string
}

// AdminUsersReport is dispatched when an Admin requests the Users report.
type AdminUsersReport struct {

	// The amount of users that are currently logged in.
	loggedInUserCount int32

	// The newest user’s ID  to be created.
	newestUserId int64

	// The username of the newest user.
	newestUsername string

	// The amount of users that have been made.
	userCount int64
}

// BanRecord contains data about a banning or unbanning which occurred, which is used by Admins to see information about Moderator actions in certain reports.
type BanRecord struct {

	// used for when a user is banned from a specific domain and not the extension as a whole.
	bannedFrom string

	// the unique id of the banrecord
	banRecordId int64

	// the id of a user that is banned.
	bannedUserId int64

	// the username of a user that is banned.
	bannedUsername string

	// Whether the user was banned (true) or unbanned (false).
	setBannedTo bool
}

// CommentVote provides the data the Front End needs to render a comment.
type Comment struct {

	// a number corresponding to a unique user ID
	userId string

	// A number corresponding to the comment’s unique ID.
	commentId int64

	// The text content of the comment.
	content string

	// An instance of CommentVote, reflecting the number of “factual” and “not factual” votes.
	factual CommentVoteDimension

	// An instance of CommentVote, reflecting the number of “funny” and “unfunny” votes.
	funny CommentVoteDimension

	// An instance of CommentVote, reflecting the number of agree and disagree votes.
	agree CommentVoteDimension

	// A boolean value, true if the comment was moderated to be hidden.
	hidden bool

	// The ID of the comment that is the parent of this comment, or 0 if it is a root-level comment.
	parent int64

	// A boolean value, true if the comment was removed.
	removed bool

	// The time the comment was posted.
	timePosted int64

	// The username of the user who posted the comment.
	username string
}

// CommentReport contains data the Front End needs to render a CommentReport, which are reports submitted by users and which Moderators can review and take action on.
type CommentReport struct {
	// If the report has been addressed
	actionTaken bool

	// The data of a comment.
	commentData Comment

	// The reason for reporting a comment.
	reasonReported string

	// The unique ID of the report.
	reportId int64

	// The unique ID of the user who reported the comment.
	reportingUserId int64

	// The name of the user that made the comment.
	reportingUsername string

	// The time that a comment was reported at.
	timeReported int64
}

// CommentVoteRecord contains data for the number of votes on a comment.
type CommentVoteDimension struct {

	// Whether the user requesting the data has already voted on the comment. It will be -1 if they have already voted down, 0 if they have not voted, and 1 if they have already voted up.
	alreadyVoted bool

	// The number of downvotes on the comment.
	downs int64

	// The number of upvotes on the comment.
	ups int64
}

// DomainModeratorRecord contains data needed by Admins to see information about DomainModerator assignments.
type DomainModeratorRecord struct {

	// the domains the moderator is allowed to moderate
	domains []string

	// When the user became aDomainModerator.
	grantedAt int64

	// The ID of the admin or GlobalModerator that promoted the user to a DomainModera tor.
	grantedBy int64

	// The username of the admin or GlobalModerator that promoted the user to a DomainModer ator.
	grantedByUsername string

	// The ID of the DomainModerator.
	grantedTo int64

	// The username of the DomainModerator.
	grantedToUsername string

	// The ID of the DomainModerators record.
	recordId int64
}

// FeedbackRecord contains data the Front End needs to render a FeedbackRecord, which is a record of a user-submitted feedback, viewed by an Admin, such as a feature request, or bug report.
type FeedbackRecord struct {

	// The text of the feedback, limited to 1000 characters.
	content string

	// Whether or not this feedback is hidden, that is to say, the admins do not want to see it again by default.
	hide bool

	// int64
	id int64

	// The time this feedback was submitted.
	submittedAt int64

	// Indicates whether this feedback is a feature request, “feature”, bug report “bug”, or
	feedbackType string

	// The ID of the user who submitted the feedback.
	userid int64

	// The username of the user who submitted the feedback.
	username string
}

// GlobalModerator record contains data needed by Admins to see information about GlobalModerator assignments.
type GlobalModeratorRecord struct {

	// When the user became a GlobalModerator.
	grantedAt int64

	// The ID of the admin that promoted the user to a GlobalModerator.
	grantedBy int64

	// The username of the admin that promoted the user to a GlobalModerator.
	grantedByUsername string

	// The ID of the GlobalModerator.
	grantedTo int64

	// The username of the GlobalModerator.
	grantedToUsername string

	// The ID of the GlobalModerators record.
	recordId int64
}

// LoginResponse is sent to the client when they successfully log in.
type LoginResponse struct {

	// The profile of a user that logged in.
	loggedInAs UserProfile
}

// ModerationRecord contains data the Front End needs to render a ModerationRecord, which is a record of a moderator action, such as hiding or removing a comment.
type ModerationRecord struct {

	// contains data the Front End needs to render a CommentReport, which are reports submitted by users and which Moderators can review and take action on.
	associatedReport CommentReport

	// The ID of the moderator's past actions.
	moderationRecordId int64

	// The id of the moderator.
	moderatorUserId int64

	// The username of the moderator.
	moderatorUsername string

	// the reason the moderator decided to take action on the comment.
	reason string

	// What they moderato set the comment's "hidden" field to
	setHiddenTo bool

	// What the moderator set the comment's "removed" field to.
	setRemovedTo bool

	// The current time that the moderator took action on a reported comment.
	timeModerated int64
}

// UserProfile contains data needed by the Front End to display a profile for a user.
type UserProfile struct {

	// The date that the user’s account was created on.
	createdOn int64

	// The server will generate a comma separated list of all domains that the user is responsible for moderating, if applicable. Otherwise, this will be an empty string.
	domainsModerating []string

	//  If the user is Admin.
	isAdmin bool

	//  If the user is DomainModerator.
	isDomainModerator bool

	// If the user is GlobalModerator.
	isGlobalModerator bool

	// The profile of the user.
	profileBlurb string

	// The ID of the user.
	userId int64

	// The name of the user.
	username string
}
