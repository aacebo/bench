CREATE TYPE user_type AS ENUM ('user', 'admin');
CREATE TABLE IF NOT EXISTS users (
    id         UUID         PRIMARY KEY,
    type       USER_TYPE    NOT NULL,
    name       VARCHAR(30)  NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS email ON users(email);
CREATE INDEX IF NOT EXISTS name ON users(name);

INSERT INTO users
(id, type, name, email, password, created_at, updated_at)
VALUES
('9ce2a810-e5c4-46b9-900f-8b4066a2a500', 'admin', 'system', 'system@bench.io', '3V4uwrZ3GGvKtm9yybj03xkmKfQ35a+W1f43AomAvRg=', NOW(), NOW());
