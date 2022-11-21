-- name: UpdateFeedbackHidden :exec
UPDATE "Feedbacks" SET hidden = $2
WHERE id = $1;

-- name: GetBanRecords :many
SELECT "BanActions".id, "BanActions".taken_by, "By".username as "taken_by_username", "BanActions".target_user, "To".username as "target_username", "BanActions".reason, "BanActions".taken_on, "BanActions".domain, "BanActions".set_banned_to FROM "BanActions" INNER JOIN "Users" as "By" on "BanActions".taken_by = "By".id INNER JOIN "Users" as "To" on "BanActions".target_user = "To".id;


-- name: GetCommentReports :many
SELECT "CR".id, "CR".reporting_user, "Users".username, "CR".comment, "CR".reason, "CR".action_taken, "CR".time_created From "CommentReports" as "CR" INNER JOIN "Users" on "CR".reporting_user = "Users".id WHERE "CR".action_taken = $1;


-- name: GetModRecordsForModerator :many 
SELECT "C"."id", "C".taken_by, "Users".username, "C".comment_id, "C".reason, "C".taken_on, "C".set_hidden_to, "C".set_removed_to, "C".associated_report FROM "CommentModerationActions" as "C" INNER JOIN "Users" ON "C".taken_by = "Users".id WHERE "C".id = $1;


-- name: GetDomainModerators :many
SELECT 
    "D"."id",
    "D"."assigned_to",
    "UTo".username as "assigned_to_username",
    "D"."assigned_at",
    "D"."assigned_by",
    "UBy".username as "assigned_by_username"
    FROM "DomainModeratorAssignments" as "D"
    INNER JOIN "Users" as "UTo" ON "D"."assigned_to".id = "UTo".id
    INNER JOIN "Users" as "UBy" ON "D"."assigned_by".id = "UBy".id
    WHERE "D"."domain" = $1 AND "D"."is_deactivation" != true;

-- name: GetGlobalModerators :many
SELECT 
    "G"."id",
    "G"."assigned_to",
    "UTo".username as "assigned_to_username",
    "G"."assigned_at",
    "G"."assigned_by",
    "UBy".username as "assigned_by_username"
    FROM "GlobalModeratorAssignments" as "G"
    INNER JOIN "Users" as "UTo" on "G".assigned_to = "UTo".id
    INNER JOIN "Users" as "UBy" on "G".assigned_by = "UBy".id
    WHERE "G".is_deactivation != true;

-- name: GetAdmins :many
SELECT 
    "G"."id",
    "G"."assigned_to",
    "UTo".username as "assigned_to_username",
    "G"."assigned_at",
    "G"."assigned_by",
    "UBy".username as "assigned_by_username"
    FROM "AdminAssignments" as "G"
    INNER JOIN "Users" as "UTo" on "G".assigned_to = "UTo".id
    INNER JOIN "Users" as "UBy" on "G".assigned_by = "UBy".id
    WHERE "G".is_deactivation != true;