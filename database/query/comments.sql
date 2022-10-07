-- name: CreateComment :exec
INSERT INTO "Comments" (
    pathid,
    author,
    content,
    response_type,
    parent
) VALUES ($1, $2, $3, $4, $5);
