-- name: GetLike :one
SELECT * FROM likes WHERE User_ID = $1 AND Quote_ID = $2;