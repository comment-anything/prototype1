-- name: CreateLog :exec
INSERT INTO "Logs" (
    user_id,
    ip,
    url 
) VALUES ($1,$2,$3);


-- name: GetLogsForDateRange :many
SELECT "L".id, "L".user_id, "U".username, "L".ip, "L".url FROM "Logs" as "L" INNER JOIN "Users" as "U" on "L".user_id = "U".id WHERE "L".at_time > $1 AND "L".at_time < $2;

-- name: GetLogsForUser :many
SELECT 
    "L"."id",
    "L"."ip",
    "L"."url",
    "L"."at_time"
    FROM "Logs" as "L"
    WHERE "L"."user_id" = $1;

-- name: GetLogsForIP :many
SELECT
    "L"."id",
    "L"."user_id",
    "U"."username",
    "L"."url",
    "L"."at_time"
    FROM "Logs" as "L"
    INNER JOIN "Users" as "U"
    ON "L"."user_id" = "Users"."id"
    WHERE "L"."ip" = $1;