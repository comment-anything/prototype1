// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: page.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const addCommentVote = `-- name: AddCommentVote :exec
INSERT INTO "VoteRecords" (
    comment_id,
    category,
    user_id,
    value
) VALUES (
    $1, $2, $3, $4
)
`

type AddCommentVoteParams struct {
	CommentID int64         `json:"comment_id"`
	Category  string        `json:"category"`
	UserID    sql.NullInt64 `json:"user_id"`
	Value     sql.NullInt64 `json:"value"`
}

func (q *Queries) AddCommentVote(ctx context.Context, arg AddCommentVoteParams) error {
	_, err := q.db.ExecContext(ctx, addCommentVote,
		arg.CommentID,
		arg.Category,
		arg.UserID,
		arg.Value,
	)
	return err
}

const getCommentVotes = `-- name: GetCommentVotes :many
select "user_id", "category", "value" From "VoteRecords" WHERE "VoteRecords"."comment_id" = $1
`

type GetCommentVotesRow struct {
	UserID   sql.NullInt64 `json:"user_id"`
	Category string        `json:"category"`
	Value    sql.NullInt64 `json:"value"`
}

func (q *Queries) GetCommentVotes(ctx context.Context, commentID int64) ([]GetCommentVotesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentVotes, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentVotesRow
	for rows.Next() {
		var i GetCommentVotesRow
		if err := rows.Scan(&i.UserID, &i.Category, &i.Value); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentsAtPath = `-- name: GetCommentsAtPath :many
select "Comments".ID, Author, Content, "Comments".Created_At, Parent, "Comments"."hidden", Removed, "Users"."username" FROM "Comments" INNER JOIN "Users" on "Comments".author = "Users"."ID" WHERE "Comments"."pathid" = $1
`

type GetCommentsAtPathRow struct {
	ID        int64         `json:"id"`
	Author    int64         `json:"author"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	Parent    sql.NullInt64 `json:"parent"`
	Hidden    sql.NullBool  `json:"hidden"`
	Removed   sql.NullBool  `json:"removed"`
	Username  string        `json:"username"`
}

func (q *Queries) GetCommentsAtPath(ctx context.Context, pathid int64) ([]GetCommentsAtPathRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsAtPath, pathid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsAtPathRow
	for rows.Next() {
		var i GetCommentsAtPathRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Content,
			&i.CreatedAt,
			&i.Parent,
			&i.Hidden,
			&i.Removed,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
