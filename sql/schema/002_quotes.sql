-- +goose Up
CREATE TABLE quotes (
    ID          UUID PRIMARY KEY,
    Created_At  TIMESTAMP    NOT NULL,
    Updated_At  TIMESTAMP    NOT NULL,
    User_ID     UUID         NOT NULL,
    Content     VARCHAR(500) NOT NULL,
    FOREIGN KEY (User_ID) REFERENCES users(ID)
);

-- -goose Down
DROP TABLE quotes;
