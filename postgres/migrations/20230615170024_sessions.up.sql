CREATE TABLE IF NOT EXISTS sessions (
    id          UUID         PRIMARY KEY,
    user_id     UUID         NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP    NOT NULL
);
