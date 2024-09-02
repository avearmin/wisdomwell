-- name: GetQuote :one
SELECT * FROM quotes WHERE ID = $1;