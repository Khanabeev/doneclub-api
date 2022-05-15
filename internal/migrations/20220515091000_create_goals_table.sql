DROP TABLE IF EXISTS doneclub.goals;
CREATE TABLE doneclub.goals
(
    id          INT UNSIGNED     NOT NULL AUTO_INCREMENT,
    user_id     INT UNSIGNED     NOT NULL,
    status      TINYINT UNSIGNED NOT NULL,
    parent_id   INT UNSIGNED,
    title       TEXT             NOT NULL,
    description LONGTEXT,
    start_date  DATETIME,
    end_date    DATETIME,
    created_at  DATETIME         NOT NULL,
    updated_at  DATETIME         NOT NULL,
    deleted_at  DATETIME,

    PRIMARY KEY (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
        REFERENCES doneclub.users (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_parent_id FOREIGN KEY (parent_id)
        REFERENCES doneclub.goals (id)
        ON DELETE SET NULL

)