
-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY "username";

-- name: GetNewestUser :one
SELECT * FROM "Users"
WHERE created_at = (SELECT MIN(created_at) FROM "Users");

-- name: GetUserCount :one
SELECT COUNT(id) FROM "Users";

