-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rating (
    user_id BIGINT  NOT NULL,
    pet_id  BIGINT  NOT NULL,
    score   FLOAT   NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rating;
-- +goose StatementEnd
