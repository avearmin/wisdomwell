-- name: GetTag :one
SELECT * FROM tags WHERE ID = $1;

-- name: PostTag :one
INSERT INTO tags (ID, Created_At, Updated_At, Name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags WHERE ID = $1;
