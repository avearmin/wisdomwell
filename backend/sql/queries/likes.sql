-- name: GetAllLikes :many
SELECT * FROM likes;

-- name: GetLike :one
SELECT * FROM likes WHERE User_ID = $1 AND Quote_ID = $2;

-- name: PostLike :one
INSERT INTO likes (User_ID, Quote_ID)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteLike :exec
DELETE FROM likes WHERE User_ID = $1 AND Quote_ID = $2;
