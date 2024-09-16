-- name: GetAllQuoteTags :many
SELECT * FROM quote_tags;

-- name: GetQuoteTag :one
SELECT * FROM quote_tags 
WHERE Quote_ID = $1 AND Tag_ID = $2;

-- name: PostQuoteTag :one
INSERT INTO quote_tags (Quote_ID, Tag_ID)
VALUES ($1, $2) RETURNING *;

-- name: DeleteQuoteTag :exec
DELETE FROM quote_tags
WHERE Quote_ID = $1 AND Tag_ID = $2;
