DROP TABLE IF EXISTS doneclub.users;
CREATE TABLE users
(
    id         INT UNSIGNED     NOT NULL AUTO_INCREMENT,
    email      VARCHAR(100)     NOT NULL UNIQUE,
    password   VARCHAR(100)     NOT NULL,
    status     TINYINT UNSIGNED NOT NULL DEFAULT 1,
    role       VARCHAR(50)      NOT NULL,
    created_at DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    PRIMARY KEY (id)
)