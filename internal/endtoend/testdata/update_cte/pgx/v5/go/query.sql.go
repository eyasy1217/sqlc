// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const updateCode = `-- name: UpdateCode :one
WITH cc AS (
            UPDATE td3.codes
            SET
                created_by = $1,
                updated_by  = $1,
                code = $2,
                hash = $3,
                is_private = false
            RETURNING hash
)
UPDATE td3.test_codes
SET
    created_by = $1,
    updated_by  = $1,
    test_id = $4,
    code_hash = cc.hash
    FROM cc
RETURNING hash, id, ts_created, ts_updated, created_by, updated_by, test_id, code_hash
`

type UpdateCodeParams struct {
	CreatedBy string
	Code      pgtype.Text
	Hash      pgtype.Text
	TestID    int32
}

type UpdateCodeRow struct {
	Hash      pgtype.Text
	ID        int32
	TsCreated pgtype.Timestamptz
	TsUpdated pgtype.Timestamptz
	CreatedBy string
	UpdatedBy string
	TestID    int32
	CodeHash  string
}

func (q *Queries) UpdateCode(ctx context.Context, arg UpdateCodeParams) (UpdateCodeRow, error) {
	row := q.db.QueryRow(ctx, updateCode,
		arg.CreatedBy,
		arg.Code,
		arg.Hash,
		arg.TestID,
	)
	var i UpdateCodeRow
	err := row.Scan(
		&i.Hash,
		&i.ID,
		&i.TsCreated,
		&i.TsUpdated,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.TestID,
		&i.CodeHash,
	)
	return i, err
}
