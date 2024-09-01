-- +goose Up
CREATE TABLE tags (
    ID UUID PRiMARY KEY,
    Created_At TIMESTAMP NOT NULL,
    Updated_At TIMESTAMP NOT NULL,
    Name VARCHAR(100) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE tags;
