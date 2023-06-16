CREATE TABLE IF NOT EXISTS problems (
    id          UUID         PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(500) NOT NULL,
    created_at  TIMESTAMP    NOT NULL
);

CREATE INDEX IF NOT EXISTS name ON problems(name);

INSERT INTO problems
(id, name, description, created_at)
VALUES
(uuid_generate_v4(), 'nbody', 'a simulation of a dynamical system of particles, usually under the influence of physical forces, such as gravity', NOW());
