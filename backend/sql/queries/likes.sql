-- name: GetAllLikes :many
SELECT * FROM likes;

-- name: GetLike :one
SELECT * FROM likes WHERE User_ID = $1 AND Quote_ID = $2;

-- name: PostLike :one
INSERT INTO likes (User_ID, Quote_ID, Created_At)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteLike :exec
DELETE FROM likes WHERE User_ID = $1 AND Quote_ID = $2;

-- name: GetAllLikesFromUser :many
SELECT * FROM likes WHERE User_ID = $1
ORDER BY Created_At DESC;
