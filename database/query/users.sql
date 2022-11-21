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

-- name: GetUserGlobalModeratorAssignments :one
SELECT "Users".id, "GlobalModeratorAssignments".assigned_at FROM "Users" INNER JOIN "GlobalModeratorAssignments" on "Users".id = "GlobalModeratorAssignments".user_id WHERE "Users".id = $1;

-- name: GetUserDomainModeratorAssignments :many
SELECT "Users".id, "DomainModeratorAssignments".assigned_at, "DomainModeratorAssignments".domain FROM "Users" INNER JOIN "DomainModeratorAssignments" on "Users".id = "DomainModeratorAssignments".user_id WHERE "Users".id = $1;

-- name: GetUserAdminAssignments :one
SELECT "Users".id, "AdminAssignments".assigned_at from "Users" INNER JOIN "AdminAssignments" on "Users".id = "AdminAssignments".user_id WHERE "Users".id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "Users"
WHERE "email" = $1 LIMIT 1;


-- name: UpdateUserPassword :exec
UPDATE "Users" SET password = $2
WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE "Users" SET email = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "Users"
WHERE id = $1;

-- name: UpdateUserBlurb :exec
UPDATE "Users" SET profile_blurb = $2
WHERE id = $1;

-- name: UpdateUserLastLogin :exec
UPDATE "Users" SET last_login = NOW()
WHERE id = $1;

-- name: UpdateUserVerification :exec
UPDATE "Users" SET is_verified = $2
WHERE id = $1;

-- name: CreateVerificationRecord :exec
INSERT INTO "VerificationCodes" (
    user_id,
    verify_code
) VALUES ($1, $2);

-- name: GetVerificationRecord :many
SELECT * FROM "VerificationCodes" WHERE user_id = $1;

-- name: DeleteVerificationRecords :exec
DELETE FROM "VerificationCodes" WHERE user_id = $1;

-- name: GetPWResetRecord :many
SELECT * FROM "PasswordResetCodes" WHERE user_id = $1;

-- name: CreatePWResetRecord :exec
INSERT INTO "PasswordResetCodes" (
    user_id,
    verify_code
) VALUES ($1,$2);

-- name: DeletePWResetRecords :exec
DELETE FROM "PasswordResetCodes" WHERE user_id = $1;