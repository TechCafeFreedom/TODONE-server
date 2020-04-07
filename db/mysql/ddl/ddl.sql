CREATE TABLE users
(
    user_id     VARCHAR(256) PRIMARY KEY,
    name        VARCHAR(128) NOT NULL,
    thumbnail   VARCHAR(256),
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     VARCHAR(256) NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_projects ON projects (user_id);

CREATE TABLE todos
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     VARCHAR(256) NOT NULL,
    title       VARCHAR(256) NOT NULL,
    content     TEXT,
    deadline    DATETIME,
    state       INT          NOT NULL DEFAULT 1,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_todos ON todos (user_id);

CREATE TABLE labels
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(256) NOT NULL,
    description VARCHAR(256),
    color       VARCHAR(256) NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_projects
(
    user_id    VARCHAR(256) NOT NULL,
    project_id INT          NOT NULL,
    role       TINYINT      NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (project_id)
        REFERENCES projects (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_users_projects ON users_projects (user_id, project_id);

CREATE TABLE projects_todos
(
    project_id INT NOT NULL,
    todo_id    INT NOT NULL,
    FOREIGN KEY (project_id)
        REFERENCES projects (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (todo_id)
        REFERENCES todos (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_projects_todos ON projects_todos (project_id, todo_id);

CREATE TABLE projects_labels
(
    project_id INT NOT NULL,
    label_id   INT NOT NULL,
    FOREIGN KEY (project_id)
        REFERENCES projects (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (label_id)
        REFERENCES labels (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_projects_labels ON projects_labels (project_id, label_id);

CREATE TABLE todos_labels
(
    todo_id INT NOT NULL,
    label_id   INT NOT NULL,
    FOREIGN KEY (todo_id)
        REFERENCES todos (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (label_id)
        REFERENCES labels (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_todos_labels ON todos_labels (todo_id, label_id);