CREATE TABLE projects
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     VARCHAR(256) NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_projects ON projects (user_id);