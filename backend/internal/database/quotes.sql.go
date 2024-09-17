// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quotes.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteQuote = `-- name: DeleteQuote :exec
DELETE FROM quotes WHERE ID = $1
`

func (q *Queries) DeleteQuote(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteQuote, id)
	return err
}

const getAllQuotes = `-- name: GetAllQuotes :many
SELECT id, created_at, updated_at, user_id, content FROM quotes
`

func (q *Queries) GetAllQuotes(ctx context.Context) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, getAllQuotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Content,
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

const getAllQuotesFromUser = `-- name: GetAllQuotesFromUser :many
SELECT id, created_at, updated_at, user_id, content FROM quotes WHERE User_ID = $1 
ORDER BY Updated_At DESC
`

func (q *Queries) GetAllQuotesFromUser(ctx context.Context, userID uuid.UUID) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, getAllQuotesFromUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Content,
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

const getQuote = `-- name: GetQuote :one
SELECT id, created_at, updated_at, user_id, content FROM quotes WHERE ID = $1
`

func (q *Queries) GetQuote(ctx context.Context, id uuid.UUID) (Quote, error) {
	row := q.db.QueryRowContext(ctx, getQuote, id)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Content,
	)
	return i, err
}

const postQuote = `-- name: PostQuote :one
INSERT INTO quotes (ID, Created_At, Updated_At, User_ID, Content)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, content
`

type PostQuoteParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	Content   string
}

func (q *Queries) PostQuote(ctx context.Context, arg PostQuoteParams) (Quote, error) {
	row := q.db.QueryRowContext(ctx, postQuote,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.Content,
	)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Content,
	)
	return i, err
}
