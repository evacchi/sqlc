// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const fn = `-- name: Fn :one
SELECT f$n()
`

func (q *Queries) Fn(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, fn)
	var f_n int32
	err := row.Scan(&f_n)
	return f_n, err
}
