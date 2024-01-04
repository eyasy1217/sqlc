// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedFilter = `-- name: SchemaScopedFilter :many
SELECT id FROM foo.bar WHERE id = $1
`

func (q *Queries) SchemaScopedFilter(ctx context.Context, id int32) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, schemaScopedFilter, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
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
