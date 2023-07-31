// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const updateSet = `-- name: UpdateSet :exec
UPDATE foo SET name = ? WHERE slug = ?
`

type UpdateSetParams struct {
	Name string
	Slug string
}

func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) error {
	_, err := q.db.ExecContext(ctx, updateSet, arg.Name, arg.Slug)
	return err
}
