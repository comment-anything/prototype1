package communication

// Ban is dispatched when a moderator or administrator bans a user.
type Ban struct {
	// The user ID that will be banned.
	UserId int64
	// Reason for the ban as provided by the moderator or administrator.
	Reason string
	// Name of domain the user will be banned from. An empty string will indicate a global ban.
	Domain string
}

// ChangeEmail is dispatched to the server when a client wants to change their email. They must supply the correct password as well.
type ChangeEmail struct {
	// The new email to associate with the client.
	NewEmail string
	// The user’s password.
	Password string
}

// ChangeFeedback is dispatched to the Server when an admin wants to remove or hide a Feedback record from being shown to them again.
type ChangeFeedback struct {
	// Whether or not to delete the Feedback
	Delete bool
	// A number representing the ID of the Feedback to change.
	FeedbackId int64
	// Whether to hide or unhide the Feedback.
	SetHiddenTo bool
}

// ChangeProfileBlurb is dispatched to the server when a client updates their profile blurb.
type ChangeProfileBlurb struct {
	// The new text which will be displayed in the user profile.
	NewBlurb string
}

// CommentReply is dispatched to the server when a logged-in user submits a reply to an existing comment or posts a new root-level comment on a page.
type CommentReply struct {
	// The id of the original comment the reply leads to.
	ReplyingTo int64
	// A comment made in a response to another comment.
	Reply string
}

// CommentVote is dispatched to the server when a logged-in user votes on a comment.
type CommentVote struct {
	// The ID of the comment the user is rating.
	VotingOn int64
	// The different types of ratings a comment. There are currently 3 planned types; “funny” for funny/unfunny comments,“factual” for factual/non factual comments, and “agree” for if the user agrees/disagrees.
	VoteType string
	// 1, if a user is voting up on a rating dimension. -1, if a user is voting down on a rating dimension, or 0, if a User is canceling their previous vote on a rating dimension.
	Value int8
}

// Feedback is dispatched to the Server when a user submits feedback on Comment
type Feedback struct {
	// Either “Bug”, “Feedback”, or “General”, depending on how the user categorizes their feedback.
	FeedbackType string
	// The text of the feedback, limited to 1,000 characters.
	Content string
}

// GetComments is dispatched to the server when a user opens the Browser Extension or when they navigate to a new page with the browser extension over. It is a request for all comments associated with the given url.
type GetComments struct {
	// The url of the current website the user is using Comment Anywhere on.
	Url string
	// Several different methods to sort the comments by, there is “time” for sorting by either latest or oldest time,  “funny” for sorting by the most or least funny comments, “factual” for comments that are the most or least factual and “agree” for the most agreed or disagreed comments.
	SortedBy string
	// Allows the user to choose if they want to sort the comments in ascending or descending order this allows for sorting by most or least relevant to the topic you're sorting by for example sorting by the oldest or newest comments.
	SortAscending bool
}

// GetUserProfile is dispatched to the server when the user needs to see a user’s profile.
type GetUserProfile struct {
	// The user ID of the target user profile.
	UserId int64
}

// Login is dispatched to the server when the client clicks “Submit” on the login form.
type Login struct {
	// The account name the user desires to login as.
	Username string
	// The password of the desired account.
	Password string
}

// Logout is dispatched to the server when the client clicks “Logout”. It does not carry any additional data.
type Logout struct {
}

// Moderate is dispatched to the server when a moderator or admin takes a moderation action on a comment.
type Moderate struct {
	// The comment Id the moderation action is being taken on.
	CommentId int64
	// The report associated with this moderation action. 0 if no report.
	AssociatedReport int64
	// The value to set hidden to.
	SetHiddenTo bool
	// The value to set removed to.
	SetRemovedTo bool
	// The reason this moderation action was taken.
	Reason string
}

// PasswordResetCode is dispatched by a user when they enter a password reset code. After a user clicks “Forgot My Password”, users may enter the code emailed to them. When they subsequently click the “submit” button, this request is dispatched to the server.
type PasswordResetCode struct {
	// The code that was sent to the user.
	Code int32
}

// PasswordReset is dispatched to the server when a password reset is requested. The client supplies the email associated with their account.
type PasswordResetRequest struct {
	// The email associated with the user to reset.
	Email string
}

// PostCommentReport is dispatched to the server when the user reports a comment.
type PostCommentReport struct {
	// The id of the comment being reported.
	CommentId int64
	// The reason the user supplied for making the report.
	Reason string
}

// Register is dispatched to the server when the client clicks “Submit” on the register form.
type Register struct {
	// The user’s chosen name for their new account.
	Username string
	// The user’s password for their new account.
	Password string
	// Retyped password for comparison.
	RetypePassword string
	// Email to be associated with the new account.
	Email string
	// Indicate if the user agreed to the terms of service.
	AgreedToTerms bool
}

// RequestValidation is dispatched to the server when the client wants a new validation code. If a client does not validate their account in a timely fashion, the validation code expires. The client may request a new validation code through their settings tab. When they do so, this entity is created and dispatched to the server.
type RequestValidation struct{}

// SetNewPass is dispatched to the Server when the user changes their password. After submitting a valid password reset code, users are prompted to set a new password. When they subsequently click “submit”, this request is dispatched to the server.
type SetNewPass struct {
	Password       string
	RetypePassword string
}

// Validate is dispatched to the server when the client inputs a validation code they received in an email to verify their account.
type Validate struct {
	Code int32
}

// ViewBans is dispatched to the server when an admin requests records of banned users.
type ViewBans struct {
	// Name of the domain. An empty string represents all domains.
	ForDomains []string
}

// ViewCommentReports is dispatched to the server when a moderator requests comment reports. It does not have any data. The server will always respond to this with all reports which have not already been moderated. If the client is a DomainModerator, the server will filter appropriately and does not require additional information from the client.
type ViewCommentReports struct {
}

// ViewFeedback is dispatched to the Server when an admin wishes to view feedback submitted by users of Comment Anywhere.
type ViewFeedback struct {
	// A time indicating the lower bound of feedback submission times.
	From int64
	// A time indicated the upper bound of feedback submission times.
	To int64
	// Either “Bug”, “Feedback”, “General”, or “All”, depending on which types of feedback the admin is interested in viewing.
	FeedbackType string
}

// ViewLogs is dispatched to the server when an admin requests access logs.
type ViewLogs struct {
	// Select a specific user. Zero returns all users.
	ForUserId int64
	// Select a specific IP address. An empty string returns all IPs.
	ForIp string
	// Name of the domain. An empty string represents all domains.
	ForDomain string
}

// ViewModRecords is dispatched to the server when an admin requests moderation records. It does not have any data. The server will always respond to this with all moderation records, sorted from newest to oldest.
type ViewModRecords struct{}

// ViewMods is dispatched to the server when an admin requests records of who has been assigned as moderators.
type ViewMods struct {
	ForDomains []string
}
