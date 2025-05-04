-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS interactions (
    user_id     BIGINT      NOT NULL,
    pet_id      BIGINT      NOT NULL,
    type        INT         NOT NULL,
    weight      FLOAT       NOT NULL,
    created_at  TIMESTAMP   NOT NULL    DEFAULT NOW(),

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS interactions;
-- +goose StatementEnd
