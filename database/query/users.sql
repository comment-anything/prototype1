-- name: CreateUser :one
INSERT INTO "Users" (
    username,
    password,
    email,
    access_level
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: GetUser :one
SELECT * FROM "Users"
WHERE "username" = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY "username";

-- name: ChangeUserPassword :exec
UPDATE "Users" SET password = $2
WHERE id = $1;

-- name: ChangeUserAccess :exec
UPDATE "Users" SET "access_level" = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;