-- +goose Up
CREATE TABLE quote_tags (
    Quote_ID UUID NOT NULL,
    Tag_ID UUID NOT NULL,
    PRIMARY KEY (Quote_ID, Tag_ID),
    FOREIGN KEY (Quote_ID) REFERENCES quotes(ID) ON DELETE CASCADE,
    FOREIGN KEY (Tag_ID) REFERENCES tags(ID) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE quote_tags;