CREATE TYPE execution_status AS ENUM ('pending', 'running', 'success', 'failure');
CREATE TABLE IF NOT EXISTS executions (
    id          UUID             PRIMARY KEY,
    solution_id UUID             NOT NULL REFERENCES solutions(id),
    test_id     UUID             NOT NULL REFERENCES tests(id),
    status      EXECUTION_STATUS NOT NULL,
    cpu_user    NUMERIC,
    cpu_sys     NUMERIC,
    mem_used    NUMERIC,
    mem_free    NUMERIC,
    stdout      TEXT,
    ouput       TEXT,
    status_at   TIMESTAMP        NOT NULL,
    created_at  TIMESTAMP        NOT NULL,
    started_at  TIMESTAMP,
    ended_at    TIMESTAMP
);
