DROP TABLE IF EXISTS doneclub.categories;
CREATE TABLE doneclub.categories
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id     INT UNSIGNED NOT NULL,
    title       TEXT         NOT NULL,
    description TEXT,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_categories_user_id FOREIGN KEY (user_id)
        REFERENCES doneclub.users (id)
        ON DELETE CASCADE
)