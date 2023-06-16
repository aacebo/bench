CREATE TABLE IF NOT EXISTS tests (
    id         UUID      PRIMARY KEY,
    problem_id UUID      NOT NULL REFERENCES problems(id),
    input      TEXT      NOT NULL,
    output     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
