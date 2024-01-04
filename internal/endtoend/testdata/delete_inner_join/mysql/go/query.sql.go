// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const removeAllAuthorsFromTheGreatGatsby = `-- name: RemoveAllAuthorsFromTheGreatGatsby :exec
DELETE author_book
FROM
  author_book
  INNER JOIN book ON book.id = author_book.book_id
WHERE
  book.title = 'The Great Gatsby'
`

func (q *Queries) RemoveAllAuthorsFromTheGreatGatsby(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, removeAllAuthorsFromTheGreatGatsby)
	return err
}
