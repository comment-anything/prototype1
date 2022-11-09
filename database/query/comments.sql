-- name: CreateComment :exec
INSERT INTO "Comments" (
    pathid,
    author,
    content,
    parent
) VALUES ($1, $2, $3, $4);
