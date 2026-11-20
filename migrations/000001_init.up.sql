CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
    id            SERIAL                 PRIMARY KEY,
    version       BIGINT       NOT NULL  DEFAULT 1,
    username      VARCHAR(100) NOT NULL  CHECK (char_length(username) BETWEEN 3 and 100),
    login         VARCHAR(30)  NOT NULL  UNIQUE CHECK (char_length(login) BETWEEN 3 AND 30),
    email         VARCHAR(254) NOT NULL  UNIQUE CHECK (
        email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
        AND
        char_length(email) BETWEEN 8 AND 254
    ),
    password_hash TEXT NOT NULL
); 

CREATE TABLE todoapp.tasks (
    id              SERIAL                 PRIMARY KEY,
    version         BIGINT        NOT NULL DEFAULT 1,
    title           VARCHAR(100)  NOT NULL CHECK (char_length(title) BETWEEN 1 and 100),
    description     VARCHAR(1000)          CHECK (char_length(description) BETWEEN 1 and 1000),
    completed       BOOLEAN       NOT NULL,
    created_at      TIMESTAMPTZ   NOT NULL,
    completed_at    TIMESTAMPTZ,

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    user_id  INTEGER       NOT NULL REFERENCES todoapp.users(id) ON DELETE CASCADE
)