-- +goose Up
CREATE TABLE users
(
    ID         UUID PRIMARY KEY,
    Created_At TIMESTAMP    NOT NULL,
    Updated_At TIMESTAMP    NOT NULL,
    Email      VARCHAR(255) NOT NULL UNIQUE,
    Name       VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;
