-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT  NOT NULL,
    name    TEXT    NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
