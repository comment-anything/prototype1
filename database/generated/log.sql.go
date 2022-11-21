// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: log.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const createLog = `-- name: CreateLog :exec
INSERT INTO "Logs" (
    user_id,
    ip,
    url 
) VALUES ($1,$2,$3)
`

type CreateLogParams struct {
	UserID sql.NullInt64  `json:"user_id"`
	Ip     sql.NullString `json:"ip"`
	Url    sql.NullString `json:"url"`
}

func (q *Queries) CreateLog(ctx context.Context, arg CreateLogParams) error {
	_, err := q.db.ExecContext(ctx, createLog, arg.UserID, arg.Ip, arg.Url)
	return err
}

const getLogsForDateRange = `-- name: GetLogsForDateRange :many
SELECT "L".id, "L".user_id, "U".username, "L".ip, "L".url FROM "Logs" as "L" INNER JOIN "Users" as "U" on "L".user_id = "U".id WHERE "L".at_time > $1 AND "L".at_time < $2
`

type GetLogsForDateRangeParams struct {
	AtTime   time.Time `json:"at_time"`
	AtTime_2 time.Time `json:"at_time_2"`
}

type GetLogsForDateRangeRow struct {
	ID       int64          `json:"id"`
	UserID   sql.NullInt64  `json:"user_id"`
	Username string         `json:"username"`
	Ip       sql.NullString `json:"ip"`
	Url      sql.NullString `json:"url"`
}

func (q *Queries) GetLogsForDateRange(ctx context.Context, arg GetLogsForDateRangeParams) ([]GetLogsForDateRangeRow, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForDateRange, arg.AtTime, arg.AtTime_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLogsForDateRangeRow
	for rows.Next() {
		var i GetLogsForDateRangeRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Username,
			&i.Ip,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsForIP = `-- name: GetLogsForIP :many
SELECT
    "L"."id",
    "L"."user_id",
    "U"."username",
    "L"."url",
    "L"."at_time"
    FROM "Logs" as "L"
    INNER JOIN "Users" as "U"
    ON "L"."user_id" = "Users"."id"
    WHERE "L"."ip" = $1
`

type GetLogsForIPRow struct {
	ID       int64          `json:"id"`
	UserID   sql.NullInt64  `json:"user_id"`
	Username string         `json:"username"`
	Url      sql.NullString `json:"url"`
	AtTime   time.Time      `json:"at_time"`
}

func (q *Queries) GetLogsForIP(ctx context.Context, ip sql.NullString) ([]GetLogsForIPRow, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForIP, ip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLogsForIPRow
	for rows.Next() {
		var i GetLogsForIPRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Username,
			&i.Url,
			&i.AtTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsForUser = `-- name: GetLogsForUser :many
SELECT 
    "L"."id",
    "L"."ip",
    "L"."url",
    "L"."at_time"
    FROM "Logs" as "L"
    WHERE "L"."user_id" = $1
`

type GetLogsForUserRow struct {
	ID     int64          `json:"id"`
	Ip     sql.NullString `json:"ip"`
	Url    sql.NullString `json:"url"`
	AtTime time.Time      `json:"at_time"`
}

func (q *Queries) GetLogsForUser(ctx context.Context, userID sql.NullInt64) ([]GetLogsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLogsForUserRow
	for rows.Next() {
		var i GetLogsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Ip,
			&i.Url,
			&i.AtTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}