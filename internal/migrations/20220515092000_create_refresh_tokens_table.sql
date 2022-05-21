DROP TABLE IF EXISTS doneclub.refresh_tokens;
CREATE TABLE doneclub.refresh_tokens
(
    id            INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id       INT UNSIGNED NOT NULL,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
)