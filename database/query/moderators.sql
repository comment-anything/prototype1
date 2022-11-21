-- name: CreateDomainModeratorAssignment :exec
INSERT INTO "DomainModeratorAssignments" (
    domain,
    assigned_to,
    assigned_by,
    is_deactivation
) VALUES (
    $1, $2, $3, $4
);

-- name: CreateGlobalModeratorAssignment :exec
INSERT INTO "GlobalModeratorAssignments" (
    assigned_to,
    assigned_by,
    is_deactivation
) VALUES (
    $1, $2, $3
);

-- name: CreateDomainBanRecord :exec
INSERT INTO "BanActions" (
    taken_by,
    target_user,
    reason,
    domain,
    set_banned_to
) VALUES ($1,$2,$3,$4,$5);

-- name: UpdateUserBanStatus :exec
UPDATE "Users" SET banned = $2
WHERE id = $1;

-- name: CreateModerationRecord :exec
INSERT INTO "CommentModerationActions" (
    taken_by,
    comment_id,
    reason,
    set_hidden_to,
    set_removed_to,
    associated_report
) VALUES ($1,$2,$3,$4,$5,$6);


-- name: UpdateCommentRemove :exec
UPDATE "Comments" SET removed = $2 WHERE id = $1;

-- name: UpdateCommentHidden :exec
UPDATE "Comments" SET hidden = $2 WHERE id = $1;