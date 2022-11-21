-- name: CreateCommentReport :exec
INSERT INTO "CommentReports" (
    reporting_user,
    comment,
    reason
) VALUES ($1, $2, $3);