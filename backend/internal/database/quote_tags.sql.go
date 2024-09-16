// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quote_tags.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteQuoteTag = `-- name: DeleteQuoteTag :exec
DELETE FROM quote_tags
WHERE Quote_ID = $1 AND Tag_ID = $2
`

type DeleteQuoteTagParams struct {
	QuoteID uuid.UUID
	TagID   uuid.UUID
}

func (q *Queries) DeleteQuoteTag(ctx context.Context, arg DeleteQuoteTagParams) error {
	_, err := q.db.ExecContext(ctx, deleteQuoteTag, arg.QuoteID, arg.TagID)
	return err
}

const getAllQuoteTags = `-- name: GetAllQuoteTags :many
SELECT quote_id, tag_id FROM quote_tags
`

func (q *Queries) GetAllQuoteTags(ctx context.Context) ([]QuoteTag, error) {
	rows, err := q.db.QueryContext(ctx, getAllQuoteTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuoteTag
	for rows.Next() {
		var i QuoteTag
		if err := rows.Scan(&i.QuoteID, &i.TagID); err != nil {
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

const getQuoteTag = `-- name: GetQuoteTag :one
SELECT quote_id, tag_id FROM quote_tags 
WHERE Quote_ID = $1 AND Tag_ID = $2
`

type GetQuoteTagParams struct {
	QuoteID uuid.UUID
	TagID   uuid.UUID
}

func (q *Queries) GetQuoteTag(ctx context.Context, arg GetQuoteTagParams) (QuoteTag, error) {
	row := q.db.QueryRowContext(ctx, getQuoteTag, arg.QuoteID, arg.TagID)
	var i QuoteTag
	err := row.Scan(&i.QuoteID, &i.TagID)
	return i, err
}

const postQuoteTag = `-- name: PostQuoteTag :one
INSERT INTO quote_tags (Quote_ID, Tag_ID)
VALUES ($1, $2) RETURNING quote_id, tag_id
`

type PostQuoteTagParams struct {
	QuoteID uuid.UUID
	TagID   uuid.UUID
}

func (q *Queries) PostQuoteTag(ctx context.Context, arg PostQuoteTagParams) (QuoteTag, error) {
	row := q.db.QueryRowContext(ctx, postQuoteTag, arg.QuoteID, arg.TagID)
	var i QuoteTag
	err := row.Scan(&i.QuoteID, &i.TagID)
	return i, err
}
