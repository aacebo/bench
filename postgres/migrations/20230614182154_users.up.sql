CREATE TABLE IF NOT EXISTS users (
    id         UUID         PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS email ON users(email);
