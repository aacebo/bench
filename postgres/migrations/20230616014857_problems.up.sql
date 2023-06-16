CREATE TABLE IF NOT EXISTS problems (
    id            UUID         PRIMARY KEY,
    name          VARCHAR(50)  NOT NULL,
    display_name  VARCHAR(50)  NOT NULL,
    description   VARCHAR(500) NOT NULL,
    created_by_id UUID         NOT NULL REFERENCES users(id),
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);

CREATE INDEX IF NOT EXISTS name ON problems(name);
