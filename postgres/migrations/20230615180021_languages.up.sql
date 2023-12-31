CREATE TABLE IF NOT EXISTS languages (
    id            UUID         PRIMARY KEY,
    icon_url      VARCHAR(255) NOT NULL,
    name          VARCHAR(50)  NOT NULL,
    version       VARCHAR(255) NOT NULL,
    is_concurrent BOOLEAN      NOT NULL,
    created_at    TIMESTAMP    NOT NULL
);

CREATE INDEX IF NOT EXISTS name ON languages(name);

INSERT INTO languages
(id, icon_url, name, version, is_concurrent, created_at)
VALUES
(uuid_generate_v4(), 'https://isocpp.org/assets/images/cpp_logo.png', 'c++', 'g++ 13.1.0', true, NOW());

INSERT INTO languages
(id, icon_url, name, version, is_concurrent, created_at)
VALUES
(uuid_generate_v4(), 'https://dev.java/assets/images/java-logo-vector.png', 'java', 'openjdk 20.0.0', true, NOW());

INSERT INTO languages
(id, icon_url, name, version, is_concurrent, created_at)
VALUES
(uuid_generate_v4(), 'https://upload.wikimedia.org/wikipedia/commons/6/6a/JavaScript-logo.png', 'javascript', 'node 18.0.0', false, NOW());
