CREATE TABLE users
(
    id         VARCHAR(256) PRIMARY KEY,
    user_id    VARCHAR(256) UNIQUE NOT NULL,
    email      VARCHAR(256) UNIQUE NOT NULL,
    password   VARCHAR(256)        NOT NULL,
    name       VARCHAR(128)        NOT NULL,
    thumbnail  VARCHAR(256)        NOT NULL,
    created_at DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME            NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users ON users (user_id, email);