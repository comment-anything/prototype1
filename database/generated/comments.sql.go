// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: comments.sql

package generated

import (
	"context"
	"database/sql"
)

const createComment = `-- name: CreateComment :exec
INSERT INTO "Comments" (
    pathid,
    author,
    content,
    parent
) VALUES ($1, $2, $3, $4)
`

type CreateCommentParams struct {
	Pathid  int64         `json:"pathid"`
	Author  int64         `json:"author"`
	Content string        `json:"content"`
	Parent  sql.NullInt64 `json:"parent"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) error {
	_, err := q.db.ExecContext(ctx, createComment,
		arg.Pathid,
		arg.Author,
		arg.Content,
		arg.Parent,
	)
	return err
}
