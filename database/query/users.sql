-- name: CreateUser :one
INSERT INTO "Users" (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetUserByUserId :one
SELECT * FROM "Users"
WHERE "id" = $1 LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM "Users"
WHERE "username" = $1 LIMIT 1;

-- name: GetUserGlobalModeratorAssignment :one
SELECT "Users".id, "GlobalModeratorAssignments".granted_at FROM "Users" INNER JOIN "GlobalModeratorAssignments" on "Users".id = "GlobalModeratorAssignments".user_id WHERE "Users".id = $1;

-- name: GetUserDomainModeratorAssignments :many
SELECT "Users".id, "DomainModeratorAssignments".granted_at, "DomainModeratorAssignments".domain FROM "Users" INNER JOIN "DomainModeratorAssignments" on "Users".id = "DomainModeratorAssignments".user_id WHERE "Users".id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "Users"
WHERE "email" = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "Users"
ORDER BY "username";

-- name: ChangeUserPassword :exec
UPDATE "Users" SET password = $2
WHERE id = $1;

-- name: ChangeUserEmail :exec
UPDATE "Users" SET email = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;