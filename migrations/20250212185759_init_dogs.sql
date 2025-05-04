-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dogs (
    id          BIGSERIAL,
    name        TEXT,
    gender      TEXT,
    breed       TEXT,
    wool_length FLOAT,
    color       TEXT,
    personality TEXT,
    size        FLOAT,
    age         FLOAT,
    weight      FLOAT,
    health      TEXT,
    experience  FLOAT,
    created_at  TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dogs;
-- +goose StatementEnd
