CREATE TABLE IF NOT EXISTS sessions (
    id          UUID      PRIMARY KEY,
    user_id     UUID      NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP NOT NULL
);

INSERT INTO sessions
(id, user_id, created_at)
VALUES
('5600e6ea-c9f1-4b86-aaa2-9f648e116660', '9ce2a810-e5c4-46b9-900f-8b4066a2a500', NOW());
