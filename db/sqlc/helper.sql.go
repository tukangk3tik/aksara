// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: helper.sql

package db

import (
	"context"
)

const clearOffices = `-- name: ClearOffices :exec
TRUNCATE TABLE offices RESTART IDENTITY CASCADE
`

func (q *Queries) ClearOffices(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearOffices)
	return err
}

const clearSchools = `-- name: ClearSchools :exec
TRUNCATE TABLE schools RESTART IDENTITY CASCADE
`

func (q *Queries) ClearSchools(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearSchools)
	return err
}

const clearUserRoles = `-- name: ClearUserRoles :exec
TRUNCATE TABLE user_roles RESTART IDENTITY CASCADE
`

func (q *Queries) ClearUserRoles(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearUserRoles)
	return err
}

const clearUsers = `-- name: ClearUsers :exec
TRUNCATE TABLE users RESTART IDENTITY CASCADE
`

func (q *Queries) ClearUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearUsers)
	return err
}
