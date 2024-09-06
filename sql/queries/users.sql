-- name: CreateUser :one
INSERT INTO users (ID, Created_At, Updated_At, Email, Name)
VALUES ($1, $2, $3, $4, $5)
returning *;

-- name: GetUser :one
SELECT * FROM users WHERE ID = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE Email = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE ID = $1;
