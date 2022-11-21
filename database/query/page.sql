
-- name: GetCommentsAtPath :many
select "Comments".ID, Author, Content, "Comments".Created_At, Parent, "Comments"."hidden", Removed, "Users"."username" FROM "Comments" INNER JOIN "Users" on "Comments".author = "Users"."ID" WHERE "Comments"."pathid" = $1;

-- name: GetCommentVotes :many
select "user_id", "category", "value" From "VoteRecords" WHERE "VoteRecords"."comment_id" = $1;

-- name: AddCommentVote :exec
INSERT INTO "VoteRecords" (
    comment_id,
    category,
    user_id,
    value
) VALUES (
    $1, $2, $3, $4
);

-- name: CreateComment :exec
INSERT INTO "Comments" (
    pathid,
    author,
    content,
    parent
) VALUES ($1, $2, $3, $4);