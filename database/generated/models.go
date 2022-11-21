// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package generated

import (
	"database/sql"
	"time"
)

type AdminAssignment struct {
	ID             int64        `json:"id"`
	AssignedTo     int64        `json:"assigned_to"`
	AssignedBy     int64        `json:"assigned_by"`
	IsDeactivation sql.NullBool `json:"is_deactivation"`
}

type BanAction struct {
	ID          int64          `json:"id"`
	TakenBy     int64          `json:"taken_by"`
	TargetUser  int64          `json:"target_user"`
	Reason      sql.NullString `json:"reason"`
	TakenOn     sql.NullTime   `json:"taken_on"`
	Domain      sql.NullString `json:"domain"`
	SetBannedTo sql.NullBool   `json:"set_banned_to"`
}

type Comment struct {
	ID        int64         `json:"id"`
	PathID    int64         `json:"path_id"`
	Author    int64         `json:"author"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	Parent    sql.NullInt64 `json:"parent"`
	Hidden    sql.NullBool  `json:"hidden"`
	Removed   sql.NullBool  `json:"removed"`
}

type CommentModerationAction struct {
	ID               int64          `json:"id"`
	TakenBy          int64          `json:"taken_by"`
	CommentId        int64          `json:"commentId"`
	Reason           sql.NullString `json:"reason"`
	TakenOn          sql.NullTime   `json:"taken_on"`
	SetHiddenTo      sql.NullBool   `json:"set_hidden_to"`
	SetRemovedTo     sql.NullBool   `json:"set_removed_to"`
	AssociatedReport sql.NullInt64  `json:"associated_report"`
}

type CommentReport struct {
	ID            int64          `json:"id"`
	ReportingUser int64          `json:"reporting_user"`
	Comment       int64          `json:"comment"`
	Reason        sql.NullString `json:"reason"`
	ActionTaken   sql.NullBool   `json:"action_taken"`
	TimeCreated   sql.NullTime   `json:"time_created"`
}

type Domain struct {
	ID string `json:"id"`
}

type DomainBan struct {
	UserID     sql.NullInt64  `json:"user_id"`
	BannedFrom sql.NullString `json:"banned_from"`
	BannedBy   sql.NullInt64  `json:"banned_by"`
	BannedAt   time.Time      `json:"banned_at"`
}

type DomainModeratorAssignment struct {
	ID             int64     `json:"id"`
	Domain         string    `json:"domain"`
	AssignedTo     int64     `json:"assigned_to"`
	AssignedAt     time.Time `json:"assigned_at"`
	AssignedBy     int64     `json:"assigned_by"`
	IsDeactivation bool      `json:"is_deactivation"`
}

type Feedback struct {
	ID          int64          `json:"id"`
	UserID      sql.NullInt64  `json:"user_id"`
	Type        sql.NullString `json:"type"`
	SubmittedAt time.Time      `json:"submitted_at"`
	Content     sql.NullString `json:"content"`
	Hidden      sql.NullBool   `json:"hidden"`
}

type GlobalModeratorAssignment struct {
	ID             int64        `json:"id"`
	AssignedTo     int64        `json:"assigned_to"`
	AssignedAt     time.Time    `json:"assigned_at"`
	AssignedBy     int64        `json:"assigned_by"`
	IsDeactivation sql.NullBool `json:"is_deactivation"`
}

type Log struct {
	ID   int64          `json:"id"`
	User sql.NullInt64  `json:"user"`
	Ip   sql.NullString `json:"ip"`
	Url  sql.NullString `json:"url"`
}

type PasswordResetCode struct {
	ID         int64          `json:"id"`
	UserID     sql.NullInt64  `json:"user_id"`
	VerifyCode sql.NullString `json:"verify_code"`
	CreatedOn  time.Time      `json:"created_on"`
}

type Path struct {
	ID     int64          `json:"id"`
	Domain sql.NullString `json:"domain"`
	Path   sql.NullString `json:"path"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	// Must be encrypted
	Password     string         `json:"password"`
	Email        string         `json:"email"`
	IsVerified   sql.NullBool   `json:"is_verified"`
	CreatedAt    time.Time      `json:"created_at"`
	LastLogin    time.Time      `json:"last_login"`
	ProfileBlurb sql.NullString `json:"profile_blurb"`
	Banned       sql.NullBool   `json:"banned"`
}

type ValidationCode struct {
	ID         int64          `json:"id"`
	UserID     sql.NullInt64  `json:"user_id"`
	VerifyCode sql.NullString `json:"verify_code"`
	CreatedOn  time.Time      `json:"created_on"`
}

type VoteRecord struct {
	CommentID int64         `json:"comment_id"`
	Category  string        `json:"category"`
	UserID    sql.NullInt64 `json:"user_id"`
	Value     sql.NullInt64 `json:"value"`
}
