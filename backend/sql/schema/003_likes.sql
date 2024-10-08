-- +goose Up
CREATE TABLE likes (
    User_ID UUID NOT NULL,
    Quote_ID UUID NOT NULL,
    Created_At TIMESTAMP NOT NULL,
    PRIMARY KEY (User_ID, Quote_ID),
    FOREIGN KEY (User_ID) REFERENCES users(ID) ON DELETE CASCADE,
    FOREIGN KEY (Quote_ID) REFERENCES quotes(ID) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE likes;
