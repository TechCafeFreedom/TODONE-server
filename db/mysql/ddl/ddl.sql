CREATE TABLE users
(
    id           INT          AUTO_INCREMENT PRIMARY KEY,
    access_token VARCHAR(256) NOT NULL,
    name         VARCHAR(128) NOT NULL,
    thumbnail    VARCHAR(256),
    created_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE boards
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     INT          NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(500),
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_boards ON boards (user_id);

CREATE TABLE kanbans
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     INT          NOT NULL,
    board_id    INT          NOT NULL,
    title       VARCHAR(256) NOT NULL,
    is_archive  TINYINT(1)   NOT NULL DEFAULT 0, -- 初期値：0（false）
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (board_id)
        REFERENCES boards (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_kanbans ON kanbans (board_id, user_id);

CREATE TABLE cards
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    user_id     INT          NOT NULL,
    kanban_id   INT          NOT NULL,
    title       VARCHAR(256) NOT NULL,
    content     TEXT,
    deadline    DATETIME,
    is_archive  TINYINT(1)   NOT NULL DEFAULT 0, -- 初期値：0（false）
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (kanban_id)
        REFERENCES kanbans (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_cards ON cards (kanban_id, user_id);

CREATE TABLE labels
(
    id          INT          AUTO_INCREMENT PRIMARY KEY,
    board_id    INT          NOT NULL,
    name        VARCHAR(256) NOT NULL,
    color       INT          NOT NULL, -- 1~12の数値
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (board_id)
        REFERENCES boards (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_labels ON labels (board_id);

CREATE TABLE users_boards
(
    user_id  INT     NOT NULL,
    board_id INT     NOT NULL,
    role     TINYINT NOT NULL,
    PRIMARY KEY (user_id, board_id),
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (board_id)
        REFERENCES boards (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE cards_labels
(
    card_id  INT NOT NULL,
    label_id INT NOT NULL,
    PRIMARY KEY (card_id, label_id),
    FOREIGN KEY (card_id)
        REFERENCES cards (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (label_id)
        REFERENCES labels (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);