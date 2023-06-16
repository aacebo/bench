CREATE TABLE IF NOT EXISTS solutions (
    id            UUID      PRIMARY KEY,
    problem_id    UUID      NOT NULL REFERENCES problems(id),
    language_id   UUID      NOT NULL REFERENCES languages(id),
    code          TEXT,
    created_by_id UUID      NOT NULL REFERENCES users(id),
    created_at    TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP NOT NULL
);
