-- name: GetAllQuotes :many
SELECT * FROM quotes;

-- name: GetQuote :one
SELECT * FROM quotes WHERE ID = $1;

-- name: PostQuote :one
INSERT INTO quotes (ID, Created_At, Updated_At, User_ID, Content)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteQuote :exec
DELETE FROM quotes WHERE ID = $1;

-- name: GetAllQuotesFromUser :many
SELECT * FROM quotes WHERE User_ID = $1 
ORDER BY Updated_At DESC;

-- name: GetRandomQuote :one
SELECT * FROM quotes
ORDER BY RANDOM()
LIMIT 1;
