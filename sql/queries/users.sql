-- name: CreateUser :one
INSERT INTO users (ID, Created_At, Updated_At, Email, Name)
VALUES ($1, $2, $3, $4, $5)
returning *;