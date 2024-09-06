-- name: GetQuote :one
SELECT * FROM quotes WHERE ID = $1;

-- name: PostQuote :one
INSERT INTO quotes (ID, Created_At, Updated_At, User_ID, Content)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteQuote :exec
DELETE FROM quotes WHERE ID = $1;