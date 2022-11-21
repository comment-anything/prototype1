-- name: CreateFeedback :exec
INSERT INTO "Feedbacks" (
    user_id,
    type,
    content
) VALUES ($1, $2, $3);

-- name: GetFeedback :many
SELECT "F".id, "F".user_id, "U".username, "F"."type", "F"."submitted_at", "F".content, "F"."hidden" FROM "Feedbacks" as "F" INNER JOIN "Users" as "U" ON "U".id = "F".user_id WHERE "F"."hidden" = $1;