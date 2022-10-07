-- name: CreateDomain :exec
INSERT INTO "Domains" (
    id
) VALUES ($1) LIMIT 1;

-- name: GetDomain :one
SELECT * FROM "Domains" 
WHERE "id" = $1 LIMIT 1;

-- name: GetPath :one
SELECT * FROM "Paths"
WHERE "domain" = $1 and "path" = $2 LIMIT 1;

-- name: CreatePath :exec
INSERT INTO "Paths" (
    domain,
    path
) VALUES ($1, $2);