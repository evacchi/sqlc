// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const columnAsOrderBy = `-- name: ColumnAsOrderBy :many
SELECT a.email AS id
FROM foo a JOIN foo b ON a.email = b.email
ORDER BY id
`

func (q *Queries) ColumnAsOrderBy(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, columnAsOrderBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
